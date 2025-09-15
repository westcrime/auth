package model

import (
	"time"
)

type Role int32

type CreateUser struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            Role
}

type User struct {
	Id        int64
	Name      string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateUser struct {
	Id   int64
	Info *UpdateUserInfo
}

type UpdateUserInfo struct {
	Name  string
	Email string
}
