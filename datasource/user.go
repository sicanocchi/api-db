package datasource

import (
	"api-db/model"
	"api-db/utils"
	"errors"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUserByID(db *gorm.DB, uid int) (model.User, error) {

	var u model.User

	if err := db.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(db *gorm.DB, username string, password string) (string, error) {

	var err error

	u := model.User{}

	err = db.Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func /*(u *User)*/ SaveUser(db *gorm.DB, u *model.User) (*model.User, error) {

	err := db.Create(&u).Error
	if err != nil {
		return &model.User{}, err
	}
	return u, nil
}

func /*(u *User)*/ BeforeSave(u *model.User) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
