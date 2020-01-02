package user

import (
	"errors"
	"pages/component"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Phone    string
	Password string
	Role     string
}

func init() {
	component.DB.AutoMigrate(&User{})
}

func (u *User) Create(user User) error {
	component.DB.Create(&user)
	if user.Username != "" {
		return nil
	}
	return errors.New("create failed")
}

func (u *User) FindUserByUsername(username string) User {
	var result User
	component.DB.Where("username = ?", username).First(&result)
	return result
}
