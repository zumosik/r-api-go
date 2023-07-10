package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zumosik/r-api-go/database"
	"github.com/zumosik/r-api-go/routes"
)

func setupRoutes(app *fiber.App) {
	// USER ROUTES
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUserById)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

}

func main() {
	app := fiber.New()     // create app
	database.ConnectToDB() // connect to db
	setupRoutes(app)       // setup all routes

	log.Fatal(app.Listen(":3000"))
}
