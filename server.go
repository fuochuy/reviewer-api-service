package main

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
	"os"
	"reviewer-api-service/db"
	"reviewer-api-service/handler"
	"reviewer-api-service/proto"
	"reviewer-api-service/store"
)

const (
	port = ":50080"
)

func main() {
	w := zerolog.ConsoleWriter{Out: os.Stderr}
	l := zerolog.New(w).With().Timestamp().Caller().Logger()
	d, err := db.New()
	if err != nil {
		err = fmt.Errorf("failed to connect database: %w", err)
		l.Fatal().Err(err).Msg("failed to connect the database")
	}
	l.Info().Str("name", d.Dialect().GetName()).
		Str("database", d.Dialect().CurrentDatabase()).
		Msg("succeeded to connect to the database")

	err = db.AutoMigrate(d)
	if err != nil {
		l.Fatal().Err(err).Msg("failed to migrate database")
	}

	ur := store.NewUserRequest(d)
	h := handler.New(&l, ur)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		l.Panic().Err(fmt.Errorf("failed to listen: %w", err))
	}

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)
	proto.RegisterUserRequestServiceServer(s, h)
	l.Info().Str("port", port).Msg("starting server")
	if err := s.Serve(lis); err != nil {
		l.Panic().Err(fmt.Errorf("failed to serve: %w", err))
	}
}
