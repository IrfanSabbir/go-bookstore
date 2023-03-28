package model

import (
	config "github.com/IrfanSabbir/go-bookstore/pkg/configs"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"" json:"name"`
	Email    string `gorm:"column:email;type:varchar(100);unique;not null"`
	Password string `gorm:"column:password;string;not null"`
	Role     string `gorm:"column:role;not null"`
	Books    []Book `gorm:"foreignkey:UserId;on_delete:cascade"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	// db.Debug().AutoMigrate(&Book{})
	// db.LogMode(true)
}
