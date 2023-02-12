package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//function to has password

func HashPassword(password string) (string,error) {
 hashedPassowrd,err:=bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
 if err != nil {
	return "", fmt.Errorf("failed to generate hash for password: %w",err)
}
return string(hashedPassowrd),nil
}

//check password method

func CheckPassword(password string,hashedPass string) error {
return bcrypt.CompareHashAndPassword([]byte(password),[]byte(hashedPass))
}