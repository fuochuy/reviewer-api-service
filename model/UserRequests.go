package model

import (
	"github.com/jinzhu/gorm"
	"reviewer-api-service/proto"
)

type UserRequests struct {
	gorm.Model
	Id             int64  `json:"id" gorm:"id"`
	UserId         int64  `json:"user_id" gorm:"user_id"`
	OrganizationId int64  `json:"organization_id" gorm:"organization_id"`
	Status         int64  `json:"status" gorm:"status"`
	Message        string `json:"message" gorm:"message"`
	//reason         string `json:"reason" gorm:"reason"`
	//amount         int64  `json:"amount" gorm:"amount"`
	//handledRole    string `json:"handled_role" gorm:"handled_role"`
}

func (userRequests *UserRequests) ProtoUserRequests() *proto.UserRequests {
	return &proto.UserRequests{
		Id:             userRequests.Id,
		UserId:         userRequests.UserId,
		OrganizationId: userRequests.OrganizationId,
		Status:         userRequests.Status,
		Message:        userRequests.Message,
		//Reason:         userRequests.reason,
		//Amount:         userRequests.amount,
		//HandledRole:    userRequests.handledRole,
	}
}
