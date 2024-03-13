package api

import (
	"net/http"

	"github.com/ThanhTien96/airbnb/internal/query"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ApiGetUserRole(db *gorm.DB) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		roles, err := query.GetUserRole(db)

		if err != nil {
			return c.JSON(http.StatusOK, jsonError(http.StatusInternalServerError, err.Error()))
		}

		return c.JSON(http.StatusOK, DataSuccessResponse("get role successfully", roles, true))
	})
}
 

// user

func ApiGetUsers(db *gorm.DB) echo.HandlerFunc {
	return echo.HandlerFunc(func (c echo.Context) error {
		users, err := query.GetAllUsers(db)

		if err != nil {
			c.JSON(http.StatusOK, jsonError(http.StatusInternalServerError, err.Error()))
		}
		return c.JSON(http.StatusOK, DataSuccessResponse("get user successfully", users, true))
	})
}
