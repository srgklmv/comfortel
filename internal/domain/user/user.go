package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	Login      string
	FirstName  string
	LastName   string
	MiddleName string
	Sex        string
	Age        uint8
	Email      string
	AvatarURL  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *User) Update(dto UpdateUserRequestDTO) {
	if dto.FirstName != "" {
		u.FirstName = dto.FirstName
	}
	if dto.LastName != "" {
		u.LastName = dto.LastName
	}
	if dto.MiddleName != "" {
		u.MiddleName = dto.MiddleName
	}
	if dto.Email != "" {
		u.Email = dto.Email
	}
	if dto.AvatarURL != "" {
		u.AvatarURL = dto.AvatarURL
	}
}
