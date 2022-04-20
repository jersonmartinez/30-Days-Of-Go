package main

import (
	"fmt"
	"log"
	"myapp/book"
	"myapp/database"

	"github.com/gofiber/fiber/v2"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connected to database")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	app.Get("/", helloWorld)

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
