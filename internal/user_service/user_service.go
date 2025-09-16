package userservice

import (
	"context"

	"github.com/westcrime/auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, createUser *model.CreateUser) (error, int64)
	Update(ctx context.Context, updateUser *model.UpdateUser) error
	Delete(ctx context.Context, user_id int64) error
	Get(ctx context.Context, user_id int64) (error, model.User)
}
