package postgres

import (
	"context"
	"errors"

	"github.com/westcrime/auth/internal/hash"
	"github.com/westcrime/auth/internal/model"
	userrepository "github.com/westcrime/auth/internal/user_repository"
	userservice "github.com/westcrime/auth/internal/user_service"
)

type userService struct {
	userRepository userrepository.UserRepository
	hashService    hash.HashService
}

func NewUserService(userRepository userrepository.UserRepository, hashService hash.HashService) userservice.UserService {
	return &userService{userRepository: userRepository, hashService: hashService}
}

func (us *userService) Create(ctx context.Context, createUser *model.CreateUser) (error, int64) {
	if createUser.Password != createUser.PasswordConfirm {
		return errors.New("Passwords are not equal"), -1
	}

	hashedPassword, err := us.hashService.HashPassword(createUser.Password)
	if err != nil {
		return err, -1
	}

	createUser.Password = hashedPassword
	createUser.PasswordConfirm = hashedPassword
	return us.userRepository.CreateUser(ctx, createUser)
}

func (us *userService) Update(ctx context.Context, updateUser *model.UpdateUser) error {
	return us.userRepository.UpdateUser(ctx, updateUser)
}

func (us *userService) Delete(ctx context.Context, id int64) error {
	return us.userRepository.DeleteUser(ctx, id)
}

func (us *userService) Get(ctx context.Context, id int64) (error, model.User) {
	return us.userRepository.GetUser(ctx, id)
}

// func GetAll(ctx context.Context, pool *pgxpool.Pool) (error, []model.User) {
// 	return query.GetUsers(ctx, pool)
// }
