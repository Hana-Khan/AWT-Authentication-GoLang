package models


import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
// Note: The gorm.Model specification adds some default properties to the Model, like id, created date, modified date, and deleted date.
// Here, the Username and Email will be unique. This means, that once we complete our application and try to register new users with the same username or email, the code won’t allow you to do it. The best part is that you don’t have to write any code specifically for this. Everything is handled by GORM.

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}