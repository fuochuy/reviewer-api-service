package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reviewer-api-service/proto"
	"strconv"
)

func (h Handler) GetUserRequestByOrgId(ctx context.Context, request *proto.GetByOrgIdRequest) (*proto.GetByOrgIdResponse, error) {
	h.logger.Info().Interface("request", request).Msg("get userRequest By Org Id")
	orgId, err := strconv.Atoi(request.GetOrganizationId())
	if err != nil {
		msg := fmt.Sprintf("cannot convert orgID (%s) into integer", request.GetOrganizationId())
		h.logger.Error().Err(err).Msg(msg)
		return nil, status.Error(codes.InvalidArgument, "invalid orgId")
	}
	userRequest, err := h.ur.GetByOrganizationId(int64(orgId))
	if err != nil {
		msg := fmt.Sprintf("userRequest (orgId=%d) not found", orgId)
		h.logger.Error().Err(err).Msg(msg)
		return nil, status.Error(codes.InvalidArgument, "invalid orgId")
	}

	return &proto.GetByOrgIdResponse{Code: 200, Message: "success", Data: userRequest.ProtoUserRequests()}, nil
}
