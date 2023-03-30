package models

import "context"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
//user request
type CreateUserRequest struct{
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
//create user response
type CreateUserResponse struct {
	ID       string  `json:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}
//login user response struct
type LoginUserResponse struct {
	Accesstoken string
	ID       string  `json:"id"`
	Username string `json:"username" db:"username"`
}
//login user response struct
type LoginUserRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
//user repository struct
type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context,email string)(*User,error)
}
// user service interface
type Service interface {
	CreateNewUser(ctx context.Context,req *CreateUserRequest)(*CreateUserResponse, error)
	Login(c context.Context,req *LoginUserRequest)(*LoginUserResponse,error)
}