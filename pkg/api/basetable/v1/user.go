package v1

import "time"

type UserRequest struct {
	Username string `json:"username" valid:"required,alphanum,stringlength(1|255)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
	Nickname string `json:"nickname" valid:"required,stringlength(1|255)"`
	Email    string `json:"email" valid:"required,email"`
	Phone    string `json:"phone" valid:"required,stringlength(11|11)"`
}

type UserLoginRequest struct {
	Username string `json:"username" valid:"required,alphanum,stringlength(1|255)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" valid:"required,stringlength(6|18)"`
	NewPassword string `json:"newPassword" valid:"required,stringlength(6|18)"`
}

type DeletedUserRequest struct {
	Ids []int `json:"ids"`
}

type ListUserRequest struct {
	Pagination
	UserRequest
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username" `
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserRequest struct {
	ID uint `json:"id" valid:"required,number"`
	UserRequest
}
