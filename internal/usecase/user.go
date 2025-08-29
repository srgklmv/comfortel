package usecase

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/srgklmv/comfortel/internal/domain/apperror"
	userDomain "github.com/srgklmv/comfortel/internal/domain/user"
	"github.com/srgklmv/comfortel/pkg/logger"
)

func (uc usecase) CreateUser(ctx context.Context, data userDomain.CreateUserRequestDTO) (any, int) {
	validationErr, err := data.Validate()
	if err != nil {
		logger.Error("data.Validate error", slog.String("error", err.Error()))
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.InternalErrorText,
		}, http.StatusInternalServerError
	}
	if validationErr != nil {
		return apperror.AppError{
			Code:    apperror.AnyIntYouWantErrorCode,
			Error:   apperror.BadRequestErrorText,
			Message: validationErr.Error(),
		}, http.StatusBadRequest
	}

	user, err := uc.userRepository.GetUserByLogin(ctx, data.Login)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Error("userRepository.GetUserByLogin error", slog.String("error", err.Error()))
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.InternalErrorText,
		}, http.StatusInternalServerError
	}
	if user.ID != uuid.Nil {
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.LoginTakenErrorText,
		}, http.StatusBadRequest
	}

	hashedPassword, err := userDomain.HashPassword(data.Password)
	if err != nil {
		logger.Error("user.HashPassword error", slog.String("error", err.Error()))
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.InternalErrorText,
		}, http.StatusInternalServerError
	}

	id, err := uc.userRepository.CreateUser(ctx, data.ToDomain(), hashedPassword)
	if err != nil {
		logger.Error("userRepository.CreateUser error", slog.String("error", err.Error()))
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.InternalErrorText,
		}, http.StatusInternalServerError
	}

	return userDomain.CreateUserResponseDTO{Created: id.String()}, http.StatusOK
}

func (uc usecase) UpdateUser(ctx context.Context, id string, data userDomain.UpdateUserRequestDTO) (any, int) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return apperror.AppError{
			Code:    apperror.AnyIntYouWantErrorCode,
			Error:   apperror.BadRequestErrorText,
			Message: "Invalid user id.",
		}, http.StatusBadRequest
	}

	validationErr, err := data.Validate()
	if err != nil {
		logger.Error("data.Validate error", slog.String("error", err.Error()))
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.InternalErrorText,
		}, http.StatusInternalServerError
	}
	if validationErr != nil {
		return apperror.AppError{
			Code:    apperror.AnyIntYouWantErrorCode,
			Error:   apperror.BadRequestErrorText,
			Message: validationErr.Error(),
		}, http.StatusBadRequest
	}

	user, err := uc.userRepository.GetUserByID(ctx, uid)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Error("userRepository.GetUserByID error", slog.String("error", err.Error()))
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.InternalErrorText,
		}, http.StatusInternalServerError
	}
	if user.ID == uuid.Nil {
		return apperror.AppError{
			Code:    apperror.AnyIntYouWantErrorCode,
			Error:   apperror.BadRequestErrorText,
			Message: "User not found.",
		}, http.StatusNotFound
	}

	user.Update(data)

	user, err = uc.userRepository.UpdateUser(ctx, user)
	if err != nil {
		logger.Error("userRepository.UpdateUser error", slog.String("error", err.Error()))
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.InternalErrorText,
		}, http.StatusInternalServerError
	}

	return userDomain.GetUserDTO{}.FromDomain(user), http.StatusOK
}

func (uc usecase) DeleteUser(ctx context.Context, id string) (any, int) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return apperror.AppError{
			Code:    apperror.AnyIntYouWantErrorCode,
			Error:   apperror.BadRequestErrorText,
			Message: "Invalid user id.",
		}, http.StatusBadRequest
	}

	user, err := uc.userRepository.GetUserByID(ctx, uid)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Error("userRepository.GetUserByID error", slog.String("error", err.Error()))
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.InternalErrorText,
		}, http.StatusInternalServerError
	}
	if user.ID == uuid.Nil {
		return apperror.AppError{
			Code:    apperror.AnyIntYouWantErrorCode,
			Error:   apperror.BadRequestErrorText,
			Message: "User not found.",
		}, http.StatusNotFound
	}

	uid, err = uc.userRepository.DeleteUser(ctx, user.ID)
	if err != nil {
		logger.Error("userRepository.DeleteUser error", slog.String("error", err.Error()))
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.InternalErrorText,
		}, http.StatusInternalServerError
	}

	return userDomain.DeleteUserResponseDTO{Deleted: uid.String()}, http.StatusOK
}

func (uc usecase) GetUsers(ctx context.Context) (any, int) {
	users, err := uc.userRepository.GetUsers(ctx)
	if err != nil {
		logger.Error("userRepository.GetUserByID error", slog.String("error", err.Error()))
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.InternalErrorText,
		}, http.StatusInternalServerError
	}

	var dtos []userDomain.GetUserDTO
	for _, user := range users {
		dtos = append(dtos, userDomain.GetUserDTO{}.FromDomain(user))
	}

	return dtos, http.StatusOK
}

func (uc usecase) GetUserByID(ctx context.Context, id string) (any, int) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return apperror.AppError{
			Code:    apperror.AnyIntYouWantErrorCode,
			Error:   apperror.BadRequestErrorText,
			Message: "Invalid user id.",
		}, http.StatusBadRequest
	}

	user, err := uc.userRepository.GetUserByID(ctx, uid)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Error("userRepository.GetUserByID error", slog.String("error", err.Error()))
		return apperror.AppError{
			Code:  apperror.AnyIntYouWantErrorCode,
			Error: apperror.InternalErrorText,
		}, http.StatusInternalServerError
	}
	if user.ID == uuid.Nil {
		return apperror.AppError{
			Code:    apperror.AnyIntYouWantErrorCode,
			Error:   apperror.BadRequestErrorText,
			Message: "User not found.",
		}, http.StatusNotFound
	}

	return userDomain.GetUserDTO{}.FromDomain(user), http.StatusOK
}
