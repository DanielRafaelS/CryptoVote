package server

import (
	"cryptovote/server/controller"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func StartServer(port string) {
	app := fiber.New()

	v1 := app.Group("api/v1")
	vote := v1.Group("cryptovote")
	vote.Get("", controller.CryptoController.GetVotes)
	vote.Post("", controller.CryptoController.SetVote)

	log.Panicln(app.Listen(fmt.Sprint(":", port)))
}
