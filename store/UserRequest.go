package store

import (
	"github.com/jinzhu/gorm"
	"reviewer-api-service/model"
)

type UserRequestRepository struct {
	db *gorm.DB
}

func NewUserRequest(db *gorm.DB) *UserRequestRepository {
	return &UserRequestRepository{
		db: db,
	}
}
func (u *UserRequestRepository) GetByOrganizationId(id int64) (*model.UserRequests, error) {
	var userRequest model.UserRequests

	err := u.db.Table("user_requests").
		Where("organization_id = ?", id).Find(&userRequest).Error
	if err != nil {
		return nil, err
	}
	return &userRequest, nil
}
