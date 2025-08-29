package user

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"time"
)

const (
	loginRegex    = "^[a-zA-Z0-9]{5,20}$"
	passwordRegex = "^[a-zA-Z0-9!&*.,#@$]{8,20}$"
	emailRegex    = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
)

type CreateUserRequestDTO struct {
	Login      string `json:"login"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
	Email      string `json:"email"`
	Sex        string `json:"sex"`
	Age        uint8  `json:"age"`
	Password   string `json:"password"`
	AvatarURL  string `json:"avatarURL"`
}

func (dto CreateUserRequestDTO) Validate() (validationError error, err error) {
	matched, err := regexp.MatchString(loginRegex, dto.Login)
	if err != nil {
		return nil, fmt.Errorf("regexp.MatchString: %w", err)
	}
	if !matched {
		validationError = errors.Join(validationError, errors.New("invalid login"))
	}

	matched, err = regexp.MatchString(passwordRegex, dto.Password)
	if err != nil {
		return nil, fmt.Errorf("regexp.MatchString: %w", err)
	}
	if !matched {
		validationError = errors.Join(validationError, errors.New("invalid password"))
	}

	matched, err = regexp.MatchString(emailRegex, dto.Email)
	if err != nil {
		return nil, fmt.Errorf("regexp.MatchString: %w", err)
	}
	if dto.Email != "" && !matched {
		validationError = errors.Join(validationError, errors.New("invalid email"))
	}

	if dto.Sex != "" && dto.Sex != "male" && dto.Sex != "female" {
		validationError = errors.Join(validationError, errors.New("invalid sex"))
	}

	if dto.Age > 150 {
		validationError = errors.Join(validationError, errors.New("age may be exaggerated a little bit"))
	}

	if len([]byte(dto.FirstName)) > 20 || len([]byte(dto.LastName)) > 20 || len([]byte(dto.MiddleName)) > 20 {
		validationError = errors.Join(validationError, errors.New("too long name"))
	}

	if dto.AvatarURL != "" {
		u, err := url.Parse(dto.AvatarURL)
		if err != nil || u.Host == "" {
			validationError = errors.Join(validationError, errors.New("invalid avatar URL"))
		}
	}

	return validationError, nil
}

func (dto CreateUserRequestDTO) ToDomain() User {
	return User{
		Login:      dto.Login,
		FirstName:  dto.FirstName,
		LastName:   dto.LastName,
		MiddleName: dto.MiddleName,
		Sex:        dto.Sex,
		Age:        dto.Age,
		AvatarURL:  dto.AvatarURL,
	}
}

type GetUserDTO struct {
	ID           string `json:"id"`
	Login        string `json:"login"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	MiddleName   string `json:"middleName,omitempty"`
	Email        string `json:"email,omitempty"`
	Sex          string `json:"sex,omitempty"`
	Age          uint8  `json:"age,omitempty"`
	AvatarURL    string `json:"avatarURL,omitempty"`
	RegisterDate string `json:"registerDate"`
}

func (dto GetUserDTO) FromDomain(u User) GetUserDTO {
	return GetUserDTO{
		ID:           u.ID.String(),
		Login:        u.Login,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		MiddleName:   u.MiddleName,
		Email:        u.Email,
		Sex:          u.Sex,
		Age:          u.Age,
		AvatarURL:    u.AvatarURL,
		RegisterDate: u.CreatedAt.Format(time.DateOnly),
	}
}

type CreateUserResponseDTO struct {
	Created string `json:"created"`
}

type UpdateUserRequestDTO struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
	Email      string `json:"email"`
	AvatarURL  string `json:"avatarURL"`
}

func (dto UpdateUserRequestDTO) Validate() (validationError error, err error) {
	// TODO: Add validations.
	return
}

type DeleteUserResponseDTO struct {
	Deleted string `json:"deleted"`
}
