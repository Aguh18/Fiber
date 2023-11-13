package utils

import "golang.org/x/crypto/bcrypt"

func HasingPassword(password string) (string, error)   {
	hasedByete, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}


	return string(hasedByete), nil
}


func CheckHasedPassword(password , hashedPassword string) bool  {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil 
}