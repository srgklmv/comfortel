package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	userDomain "github.com/srgklmv/comfortel/internal/domain/user"
)

func (r repository) CreateUser(ctx context.Context, data userDomain.User, hashedPassword string) (uuid.UUID, error) {
	tx, err := getTxFromContext(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("getTxFromContext: %w", err)
	}

	isActive := true
	entity := userDomain.EntityFromDomain(data)
	entity.IsActive = &isActive

	var id uuid.UUID

	err = tx.QueryRowContext(
		ctx,
		`insert into "user" (login, email, first_name, last_name, middle_name, sex, age, avatar_url, is_active, password)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		returning id;`,
		entity.Login, entity.Email, entity.FirstName, entity.LastName, entity.MiddleName, entity.Sex, entity.Age, entity.AvatarURL, entity.IsActive, hashedPassword,
	).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("queryRowContext: %w", err)
	}

	return id, nil
}

func (r repository) UpdateUser(ctx context.Context, user userDomain.User) (userDomain.User, error) {
	tx, err := getTxFromContext(ctx)
	if err != nil {
		return user, fmt.Errorf("getTxFromContext: %w", err)
	}

	entity := userDomain.EntityFromDomain(user)

	var fields []string
	var args []any
	if entity.Email != nil {
		fields = append(fields, fmt.Sprintf("email = $%d", len(fields)+1))
		args = append(args, &entity.Email)
	}
	if entity.FirstName != nil {
		fields = append(fields, fmt.Sprintf("first_name = $%d", len(fields)+1))
		args = append(args, &entity.FirstName)
	}
	if entity.LastName != nil {
		fields = append(fields, fmt.Sprintf("last_name = $%d", len(fields)+1))
		args = append(args, &entity.LastName)
	}
	if entity.MiddleName != nil {
		fields = append(fields, fmt.Sprintf("middle_name = $%d", len(fields)+1))
		args = append(args, &entity.MiddleName)
	}
	if entity.AvatarURL != nil {
		fields = append(fields, fmt.Sprintf("avatar_url = $%d", len(fields)+1))
		args = append(args, &entity.AvatarURL)
	}

	if len(fields) == 0 {
		return user, errors.New("no fields passed")
	}

	set := strings.Join(fields, ", ")
	query := []string{
		`update "user" set`,
		set,
		fmt.Sprintf("where id = $%d", len(fields)+1),
		`returning id, login, email, first_name, last_name, middle_name, sex, age, avatar_url, is_active, created_at;`,
	}
	args = append(args, user.ID)
	q := strings.Join(query, " ")

	err = tx.QueryRowContext(
		ctx,
		q,
		args...,
	).Scan(
		&entity.ID,
		&entity.Login,
		&entity.Email,
		&entity.FirstName,
		&entity.LastName,
		&entity.MiddleName,
		&entity.Sex,
		&entity.Age,
		&entity.AvatarURL,
		&entity.IsActive,
		&entity.CreatedAt,
	)
	if err != nil {
		return user, fmt.Errorf("queryRowContext: %w", err)
	}

	return entity.ToDomain(), nil
}

func (r repository) GetUserByLogin(ctx context.Context, login string) (userDomain.User, error) {
	tx, err := getTxFromContext(ctx)
	if err != nil {
		return userDomain.User{}, fmt.Errorf("getTxFromContext: %w", err)
	}

	var e userDomain.Entity

	err = tx.QueryRowContext(
		ctx,
		`select id, login, email, first_name, last_name, middle_name, sex, age, created_at, updated_at, avatar_url
		from "user"
		where login = $1;`,
		login,
	).Scan(&e.ID, &e.Login, &e.Email, &e.FirstName, &e.LastName, &e.MiddleName, &e.Sex, &e.Age, &e.CreatedAt, &e.UpdatedAt, &e.AvatarURL)
	if err != nil {
		return userDomain.User{}, fmt.Errorf("queryRowContext: %w", err)
	}

	return e.ToDomain(), nil
}

func (r repository) DeleteUser(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	tx, err := getTxFromContext(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("getTxFromContext: %w", err)
	}

	err = tx.QueryRowContext(
		ctx,
		`delete
		from "user"
		where login = $1;`,
		id,
	).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("queryRowContext: %w", err)
	}

	return id, nil
}

func (r repository) GetUserByID(ctx context.Context, id uuid.UUID) (userDomain.User, error) {
	tx, err := getTxFromContext(ctx)
	if err != nil {
		return userDomain.User{}, fmt.Errorf("getTxFromContext: %w", err)
	}

	var e userDomain.Entity

	err = tx.QueryRowContext(
		ctx,
		`select id, login, email, first_name, last_name, middle_name, sex, age, created_at, updated_at, avatar_url
		from "user"
		where id = $1;`,
		id,
	).Scan(&e.ID, &e.Login, &e.Email, &e.FirstName, &e.LastName, &e.MiddleName, &e.Sex, &e.Age, &e.CreatedAt, &e.UpdatedAt, &e.AvatarURL)
	if err != nil {
		return userDomain.User{}, fmt.Errorf("queryRowContext: %w", err)
	}

	return e.ToDomain(), nil
}

func (r repository) GetUsers(ctx context.Context) ([]userDomain.User, error) {
	tx, err := getTxFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getTxFromContext: %w", err)
	}

	var users []userDomain.User

	rows, err := tx.QueryContext(
		ctx,
		`select id, login, email, first_name, last_name, middle_name, sex, age, created_at, updated_at, avatar_url
		from "user";`,
	)
	if err != nil {
		return nil, fmt.Errorf("queryRowContext: %w", err)
	}

	for rows.Next() {
		var e userDomain.Entity
		err = rows.Scan(&e.ID, &e.Login, &e.Email, &e.FirstName, &e.LastName, &e.MiddleName, &e.Sex, &e.Age, &e.CreatedAt, &e.UpdatedAt, &e.AvatarURL)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		users = append(users, e.ToDomain())
	}

	return users, nil
}
