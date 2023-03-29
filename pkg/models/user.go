package model

import (
	"fmt"

	config "github.com/IrfanSabbir/go-bookstore/pkg/configs"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `gorm:"" json:"name"`
	Email    string `gorm:"column:email;type:varchar(100);unique;not null"`
	Password string `gorm:"column:password;string;not null"`
	Role     string `gorm:"column:role;not null"`
	Books    []Book `gorm:"foreignkey:UserId;on_delete:cascade"`
}

type AuthUserResponse struct {
	user    User
	message string
	token   string
}

func init() {
	config.Connect()
	db = config.GetDB()
	// db.Debug().AutoMigrate(&Book{})
	// db.LogMode(true)
}

func GetUserByEmail(email string) User {
	var user User
	db.Where("email = ?", email).Find(&user)
	return user
}

func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHsh(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (user User) RegisterUser() User {
	existedUser := GetUserByEmail(user.Email)
	if existedUser.Email != "" {
		fmt.Println(existedUser)
		return user
	}
	hashedPassword, err := HashedPassword(user.Password)
	if err != nil {
		return user
	}
	user.Password = hashedPassword
	if user.Role == "" {
		user.Role = "Member"
	}
	db.Create(&user)
	db.NewRecord(user)
	return user
}

func Login(email string, password string) User {
	existedUser := GetUserByEmail(email)
	fmt.Println(existedUser)
	if existedUser.Email == "" {
		fmt.Println("Email not valid")
		// return AuthUserResponse{user: existedUser, message: "Email not valid", token: ""}
		return User{}
	}

	isMatched := CheckPasswordHsh(existedUser.Password, password)
	if isMatched == false {
		fmt.Println("PAssword not valid")
		// return AuthUserResponse{user: existedUser, message: "PAssword not valid", token: ""}
		return User{}
	}
	var response AuthUserResponse
	response.user = existedUser
	response.message = "Auth Successfull"
	response.token = "this_is_temp_token"
	return existedUser
}
