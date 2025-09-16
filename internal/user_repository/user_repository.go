package userrepository

import (
	"context"

	"github.com/westcrime/auth/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, createUser *model.CreateUser) (error, int64)
	UpdateUser(ctx context.Context, updateUser *model.UpdateUser) error
	DeleteUser(ctx context.Context, user_id int64) error
	GetUser(ctx context.Context, user_id int64) (error, model.User)
}
