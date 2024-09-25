package apis

import (
	"log"

	"github.com/gofiber/fiber/v2"

	databases "tryGo_two/database"
)

func ApisTest() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		usr := databases.GetUsers()
		if usr == nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving users")
		}
		return c.Status(fiber.StatusOK).SendString(string(usr))
	})

	//start server
	log.Fatal(app.Listen(":3000"))
}
