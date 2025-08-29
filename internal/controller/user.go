package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srgklmv/comfortel/internal/domain/apperror"
	userDomain "github.com/srgklmv/comfortel/internal/domain/user"
)

func (c controller) CreateUser(gc *gin.Context) {
	var body userDomain.CreateUserRequestDTO
	err := gc.ShouldBindJSON(&body)
	if err != nil {
		gc.JSON(http.StatusBadRequest, apperror.AppError{
			Code:    apperror.AnyIntYouWantErrorCode,
			Error:   apperror.BadRequestErrorText,
			Message: "Request body invalid.",
		})
		return
	}

	response, status := c.userUsecase.CreateUser(gc, body)

	gc.JSON(status, response)
}

func (c controller) GetUser(gc *gin.Context) {
	id := gc.Param("id")

	response, status := c.userUsecase.GetUserByID(gc, id)

	gc.JSON(status, response)
}

func (c controller) GetUsers(gc *gin.Context) {
	response, status := c.userUsecase.GetUsers(gc)

	gc.JSON(status, response)
}

func (c controller) UpdateUser(gc *gin.Context) {
	id := gc.Param("id")

	var dto userDomain.UpdateUserRequestDTO
	err := gc.ShouldBindJSON(&dto)
	if err != nil {
		gc.JSON(http.StatusBadRequest, apperror.AppError{
			Code:    apperror.AnyIntYouWantErrorCode,
			Error:   apperror.BadRequestErrorText,
			Message: "Request body invalid.",
		})
		return
	}

	response, status := c.userUsecase.UpdateUser(gc, id, dto)

	gc.JSON(status, response)
}

func (c controller) DeleteUser(gc *gin.Context) {
	id := gc.Param("id")

	response, status := c.userUsecase.DeleteUser(gc, id)

	gc.JSON(status, response)
}
