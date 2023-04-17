package model

import (
	"fmt"

	config "github.com/IrfanSabbir/go-bookstore/pkg/configs"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	// ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"size:255;not null" json:"name"`
	Author      string `gorm:"size:255;not null" json:"author"`
	Publication string `gorm:"size:255;not null" json:"publication"`
	// CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	// UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	User   User   `json:"user"`
	UserId uint64 `gorm:"not null" json:"user_id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	// db.Debug().AutoMigrate(&Book{})
	// db.LogMode(true)
}

func GetAllBooks() []Book {
	var Books []Book
	db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, role")
	}).Find(&Books)
	return Books
}

func GetBookById(id int64) Book {
	var book Book
	db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, role")
	}).Where("id = ?", id).Find(&book)
	return book
}

func (createBook Book) CreateBook(user_id int64) Book {
	user, err := GetUserById(user_id)
	if err != nil {

		return Book{}
	}
	bookRecord := Book{
		Name:        createBook.Name,
		Author:      createBook.Author,
		Publication: createBook.Publication,
		User:        user,
	}
	db.NewRecord(bookRecord)
	db.Create(&bookRecord)
	return bookRecord
}

func DeleteBook(id int64, user_id int64) (Book, error) {
	book := GetBookById(id)

	if book.UserId == uint64(user_id) {
		db.Where("id = ?", id).Delete(&book)
		return book, nil
	} else {
		return book, fmt.Errorf("You are not authorized to perform this query")
	}
}

func (updatedItem Book) UpadteBook(id int64, user_id int64) (Book, error) {
	var currentBook Book
	db.Where("id = ?", id).Find(&currentBook)

	if currentBook.UserId != uint64(user_id) {
		return updatedItem, fmt.Errorf("You are not authorized to perform this query")
	}
	if updatedItem.Name != "" {
		currentBook.Name = updatedItem.Name
	}

	if updatedItem.Publication != "" {
		currentBook.Publication = updatedItem.Publication
	}
	if updatedItem.Author != "" {
		currentBook.Author = updatedItem.Author
	}
	// db.Model(&currentBook).Update("name", "I am going home")
	db.Save(&currentBook)
	return currentBook, nil
}
