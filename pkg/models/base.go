package model

import (
	config "github.com/IrfanSabbir/go-bookstore/pkg/configs"
)

func init() {
	config.Connect()
	db = config.GetDB()
	db.Debug().AutoMigrate(&Book{}, &User{})
	db.LogMode(true)
}
