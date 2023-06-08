package webserver

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/koksmat-com/koksmat/model"
)

type UserResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

func Serve() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/v1/:name", func(c *fiber.Ctx) error {

		switch c.Params("name") {

		case "sharedmailboxes":
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var sharedMailboxes []model.SharedMailbox
			defer cancel()

			results, err := model.GetSharedMailboxes()
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
			}
			defer results.Close(ctx)
			for results.Next(ctx) {
				var sharedMailbox model.SharedMailbox
				if err = results.Decode(&sharedMailbox); err != nil {
					return c.Status(http.StatusInternalServerError).JSON(UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
				}

				sharedMailboxes = append(sharedMailboxes, sharedMailbox)
			}
			return c.Status(http.StatusOK).JSON(
				UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": sharedMailboxes}},
			)
		default:

			return c.Status(404).SendString("Where is john?")

		}

	})

	log.Fatal(app.Listen(":3000"))
}
