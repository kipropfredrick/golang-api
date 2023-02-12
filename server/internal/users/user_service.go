package users

import (
	"context"
	"fmt"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//to be moved to env file
const(
	secret="sasasassa"
)
//type struct thataccept repository as arguement
type UserService struct {
	Repository
	timeout time.Duration
}

//function constructor that usersivercie struct as arguement and
func NewUserService(repository Repository) Service {
	// return pointer to user service struct
    return &UserService{
		repository,
		time.Duration(2)*time.Second,
	}
}
//jwt claims struct
type MyjwtClaims struct {
	ID string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
//function to create user 
func (s *UserService) CreateNewUser(c context.Context,req *CreateUserRequest)(*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	fmt.Println("hased",hashedPassword)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res:=&CreateUserResponse{
		ID:strconv.Itoa(int(r.ID)),//type cast to string
		Email: r.Email,
		Username: req.Username,
	}

	return res, nil

}
//login user
func (s UserService) Login(c context.Context,req *LoginUserRequest)(*LoginUserResponse,error){
ctx ,cancel:=context.WithTimeout(c,s.timeout)
defer cancel()
u,err :=s.Repository.GetByEmail(ctx,req.Email)
if err != nil {
	return &LoginUserResponse{},err
}
hashedPassword, _ := util.HashPassword(req.Password)
fmt.Println("u",u,hashedPassword)
err= util.CheckPassword(req.Password,u.Password)
if err != nil {
	return &LoginUserResponse{}, err
}
//generate jwt if password matches the database
  token :=jwt.NewWithClaims(jwt.SigningMethodHS256,MyjwtClaims{
	ID: strconv.Itoa(int(u.ID)),
	Username: u.Username,
	RegisteredClaims: jwt.RegisteredClaims{
		Issuer: strconv.Itoa(int(u.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
	
	},
  })
  signedstring,err :=token.SignedString([]byte(secret))
  if err != nil {
	return &LoginUserResponse{}, err
  }
  return &LoginUserResponse{accesstoken: signedstring,ID:strconv.Itoa(int(u.ID)),Username: u.Username}, nil
}