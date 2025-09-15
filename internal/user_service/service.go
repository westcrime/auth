package userservice

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/westcrime/auth/internal/hash"
	"github.com/westcrime/auth/internal/model"
	"github.com/westcrime/auth/internal/query"
)

func Create(ctx context.Context, pool *pgxpool.Pool, createUser *model.CreateUser, hashService hash.HashService) (error, int64) {
	if createUser.Password != createUser.PasswordConfirm {
		return errors.New("Passwords are not equal"), -1
	}

	hashedPassword, err := hashService.HashPassword(createUser.Password)
	if err != nil {
		return err, -1
	}

	createUser.Password = hashedPassword
	createUser.PasswordConfirm = hashedPassword
	return query.CreateUser(ctx, pool, createUser)
}

func Update(ctx context.Context, pool *pgxpool.Pool, updateUser *model.UpdateUser) error {
	return query.UpdateUser(ctx, pool, updateUser)
}

func Delete(ctx context.Context, pool *pgxpool.Pool, id int64) error {
	return query.DeleteUser(ctx, pool, id)
}

func Get(ctx context.Context, pool *pgxpool.Pool, id int64) (error, model.User) {
	return query.GetUser(ctx, pool, id)
}

// func GetAll(ctx context.Context, pool *pgxpool.Pool) (error, []model.User) {
// 	return query.GetUsers(ctx, pool)
// }
