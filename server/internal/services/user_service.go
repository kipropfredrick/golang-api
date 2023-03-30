package services

import (
	"context"
	"fmt"
	"server/util"
	"strconv"
	"time"
	"server/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

//to be moved to env file
const(
	secret="sasasassa"
)
//type struct thataccept repository as arguement
type UserService struct {
	models.Repository
	timeout time.Duration
}

//function constructor that usersivercie struct as arguement and
func NewUserService(repository models.Repository) models.Service {
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
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
//function to create user 
func (s *UserService) CreateNewUser(c context.Context,req *models.CreateUserRequest)(*models.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	fmt.Println("hased",hashedPassword)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res:=&models.CreateUserResponse{
		ID:strconv.Itoa(int(r.ID)),//type cast to string
		Email: r.Email,
		Username: req.Username,
	}

	return res, nil

}
//login user
func (s UserService) Login(c context.Context,req *models.LoginUserRequest)(*models.LoginUserResponse,error){
ctx ,cancel:=context.WithTimeout(c,s.timeout)
defer cancel()
u,err :=s.Repository.GetByEmail(ctx,req.Email)
if err != nil {
	return &models.LoginUserResponse{},err
}

hashedPassword, _ := util.HashPassword(req.Password)
fmt.Println("u",u,hashedPassword)
err= util.CheckPassword(req.Password,u.Password)
if err != nil {
	return &models.LoginUserResponse{}, err
}
//generate jwt if password matches the database
  token :=jwt.NewWithClaims(jwt.SigningMethodHS256,MyjwtClaims{
	ID: strconv.Itoa(int(u.ID)),
	Username: u.Username,
	// Name: u.Name,
	// email: u.Email,
	RegisteredClaims: jwt.RegisteredClaims{
		Issuer: strconv.Itoa(int(u.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
	
	},
  })
  signedstring,err :=token.SignedString([]byte(secret))
  if err != nil {
	return &models.LoginUserResponse{}, err
  }
  return &models.LoginUserResponse{ Accesstoken:signedstring,ID:strconv.Itoa(int(u.ID)),Username: u.Username}, nil
}