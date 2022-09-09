package handler

import (
	"github.com/rs/zerolog"
	"reviewer-api-service/store"
)

type Handler struct {
	logger *zerolog.Logger
	ur     *store.UserRequestRepository
}

// New returns a new handler with logger and database
func New(l *zerolog.Logger, ur *store.UserRequestRepository) *Handler {
	return &Handler{logger: l, ur: ur}
}
