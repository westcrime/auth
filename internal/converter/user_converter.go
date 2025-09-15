package converter

import (
	"github.com/westcrime/auth/internal/model"
	desc "github.com/westcrime/auth/pkg/user_v1"
)

func ToCreateUserModelFromCreateRequestProto(req *desc.CreateRequest) *model.CreateUser {
	return &model.CreateUser{
		Name:            req.Info.Name,
		Email:           req.Info.Email,
		Password:        req.Info.Password,
		PasswordConfirm: req.Info.PasswordConfirm,
		Role:            model.Role(req.Info.Role),
	}
}

func ToUpdateUserModelFromUpdateRequestProto(req *desc.UpdateRequest) *model.UpdateUser {
	return &model.UpdateUser{
		Id: req.Id,
		Info: &model.UpdateUserInfo{
			Email: req.Info.Email.Value,
			Name:  req.Info.Name.Value,
		},
	}
}

func ToUserModelFromUserProto(req *desc.User) *model.User {
	return &model.User{
		Id:        req.Id,
		Email:     req.Info.Email,
		Name:      req.Info.Name,
		Role:      model.Role(req.Info.Role),
		CreatedAt: req.CreatedAt.AsTime(),
		UpdatedAt: req.UpdatedAt.AsTime(),
	}
}
