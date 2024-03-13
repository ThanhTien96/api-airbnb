package utils

import (
	"fmt"
	"log"

	"github.com/ThanhTien96/airbnb/models"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

const (
	USER_LIMIT = 100
)

func CreateUsers(db *gorm.DB, limit int) {
	log.Println("Creating Movies...")

	for i := 0; i <  limit; i++ {
		user := models.UserBase{
			CreateUserRequest: models.CreateUserRequest{
				UserName: gofakeit.Name(),
				Password: gofakeit.Password(true, true, true, true, true, 4),
				Email: gofakeit.Email(),
				PhoneNumber: gofakeit.Phone(),
				RoleID: 28,
				Avatar: gofakeit.ImageURL(640, 480),
			},
		}
		_ = db.Create(&user)
	}	
}

func CreateUserRole(db *gorm.DB) {
	roleName := []string{
		"user",
		"admin",
		"manager",
	}
	for _,i := range(roleName) {
		fmt.Println(i)
		role := models.UserRole{
			RoleName: i,
		}
		_ = db.Create(&role)
	}	
}


func CleanData(db *gorm.DB) {
	var users []models.User
	log.Println("Deleting Users")
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&users)
	log.Println("Reset sequence of users")
	db.Exec("ALTER SEQUENCE users_user_id_seq RESTART WITH 1;")

	var userRoles []models.UserRole
	log.Println("Deleting userRoles")
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&userRoles)
	log.Println("Reset sequence of user roles")
	db.Exec("ALTER SEQUENCE userRole_role_id_seq RESTART WITH 1;")
}
