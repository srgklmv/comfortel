package usecase

import (
	"context"

	"github.com/google/uuid"
	userDomain "github.com/srgklmv/comfortel/internal/domain/user"
)

type repository interface {
	userRepository
}

type userRepository interface {
	GetUserByLogin(ctx context.Context, login string) (userDomain.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (userDomain.User, error)
	GetUsers(ctx context.Context) ([]userDomain.User, error)
	CreateUser(ctx context.Context, data userDomain.User, hashedPassword string) (uuid.UUID, error)
	UpdateUser(ctx context.Context, data userDomain.User) (userDomain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
}

type usecase struct {
	userRepository userRepository
}

func New(repository repository) *usecase {
	return &usecase{
		userRepository: repository,
	}
}
