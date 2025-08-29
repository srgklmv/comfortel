package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/srgklmv/comfortel/pkg/utils/pointer"
)

type Entity struct {
	ID         *uuid.UUID
	Login      *string
	FirstName  *string
	LastName   *string
	MiddleName *string
	Sex        *string
	Age        *uint8
	Email      *string
	AvatarURL  *string
	IsActive   *bool
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func EntityFromDomain(u User) Entity {
	var e Entity

	if u.ID != uuid.Nil {
		e.ID = &u.ID
	}
	if u.Login != "" {
		e.Login = &u.Login
	}
	if u.Email != "" {
		e.Email = &u.Email
	}
	if u.FirstName != "" {
		e.FirstName = &u.FirstName
	}
	if u.LastName != "" {
		e.LastName = &u.LastName
	}
	if u.MiddleName != "" {
		e.MiddleName = &u.MiddleName
	}
	if u.Sex != "" {
		e.Sex = &u.Sex
	}
	if u.Age != 0 {
		e.Age = &u.Age
	}
	if u.AvatarURL != "" {
		e.AvatarURL = &u.AvatarURL
	}

	return e
}

func (e Entity) ToDomain() User {
	return User{
		ID:         pointer.ParsePointer(e.ID),
		Login:      pointer.ParsePointer(e.Login),
		FirstName:  pointer.ParsePointer(e.FirstName),
		LastName:   pointer.ParsePointer(e.LastName),
		MiddleName: pointer.ParsePointer(e.MiddleName),
		Sex:        pointer.ParsePointer(e.Sex),
		Age:        pointer.ParsePointer(e.Age),
		Email:      pointer.ParsePointer(e.Email),
		AvatarURL:  pointer.ParsePointer(e.AvatarURL),
		CreatedAt:  pointer.ParsePointer(e.CreatedAt),
		UpdatedAt:  pointer.ParsePointer(e.UpdatedAt),
	}
}
