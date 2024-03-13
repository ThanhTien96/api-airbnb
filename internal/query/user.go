package query

import (
	"github.com/ThanhTien96/airbnb/models"
	"gorm.io/gorm"
)

func GetUserRole(db *gorm.DB) ([]*models.UserRole, error) {
	var role []*models.UserRole

	result := db.Table(models.RoleTableName).Find(&role)

	if result.Error != nil {
		return nil, result.Error
	}

	return role, nil
}


// user

func GetAllUsers(db *gorm.DB) ([]*models.UserBase, error) {
	var users []*models.UserBase

	result := db.Table(models.UserTableName).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}