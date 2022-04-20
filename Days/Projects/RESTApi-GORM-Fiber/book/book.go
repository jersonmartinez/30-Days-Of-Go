package book

import (
	"myapp/database"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {

	db := database.DBConn

	var books []Book
	db.Find(&books)

	return c.JSON(books)

	// return c.SendString("All books")
}

func GetBook(c *fiber.Ctx) error {

	db := database.DBConn
	id := c.Params("id")

	var book Book
	db.First(&book, id)

	return c.JSON(book)

	// return c.SendString("A single book")
}

func NewBook(c *fiber.Ctx) error {

	db := database.DBConn
	book := Book{}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(503).SendString("Failed to parse body")
	}

	// book.Title = "1984"
	// book.Author = "George Orwell"
	// book.Rating = 5

	db.Create(&book)

	return c.JSON(book)

	// return c.SendString("Add a new book")
}

func DeleteBook(c *fiber.Ctx) error {

	db := database.DBConn
	id := c.Params("id")

	var book Book
	db.First(&book, id)

	if book.Title == "" {
		return c.SendString("No book found with given ID")
	}

	db.Delete(&book)

	return c.SendString("Book deleted: " + id + " => " + book.Title)

	// return c.SendString("Delete a book")
}
