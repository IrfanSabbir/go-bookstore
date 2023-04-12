package model

import (
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
	db.Find(&Books)
	return Books
}

func GetBookById(id int64) Book {
	var book Book
	db.Where("id = ?", id).Find(&book)
	return book
}

func (createBook Book) CreateBook() Book {
	db.NewRecord(createBook)
	db.Create(&createBook)
	return createBook
}

func DeleteBook(id int64) Book {
	book := GetBookById(id)
	db.Where("id = ?", id).Delete(&book)
	return book
}

func (updatedItem Book) UpadteBook(id int64) Book {
	currentBook := GetBookById(id)
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
	return currentBook
}
