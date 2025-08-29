package controller

import (
	"context"

	"github.com/srgklmv/comfortel/internal/domain/user"
)

type controller struct {
	userUsecase userUsecase
}

type usecase interface {
	userUsecase
}

type userUsecase interface {
	CreateUser(ctx context.Context, data user.CreateUserRequestDTO) (any, int)
	GetUserByID(ctx context.Context, id string) (any, int)
	GetUsers(ctx context.Context) (any, int)
	UpdateUser(ctx context.Context, id string, data user.UpdateUserRequestDTO) (any, int)
	DeleteUser(ctx context.Context, id string) (any, int)
}

func New(uc usecase) *controller {
	return &controller{
		userUsecase: uc,
	}
}
