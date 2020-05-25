package main

import (
	"github.com/fabiopimentel/Agenda/pkg/controller"
	"github.com/gofiber/fiber"
)


func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	a := controller.AgendaController{}
	a.Init()

	app.Get("/api/v1/contact", a.GetContacts)
	app.Get("/api/v1/contact/:id", a.GetContact)
	app.Post("/api/v1/contact", a.NewContact)
	app.Delete("/api/v1/contact/:id", a.DeleteContact)
}
func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(5000)
}
