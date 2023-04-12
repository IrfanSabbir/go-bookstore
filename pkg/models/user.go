package model

import (
	"fmt"

	config "github.com/IrfanSabbir/go-bookstore/pkg/configs"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	// ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:255;unique;not null" json:"email"`
	Password string `gorm:"string;not null" json:"password"`
	Role     string `gorm:"not null" json:"role"`
	// CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	// UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Books []Book `gorm:"foreignkey:UserId;constraint:OnDelete:CASCADE" json:"books"`
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
		return User{}
	}

	isMatched := CheckPasswordHsh(existedUser.Password, password)
	if isMatched == false {
		fmt.Println("PAssword not valid")
		return User{}
	}
	return existedUser
}
