package routes

import (
	"eunoia/internal/chat"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Register(app *fiber.App) {
	app.Post("/create",chat.CreateRoom)
	app.Post("/join", chat.JoinRoom)
	app.Get("/ws/:roomID", websocket.New(chat.HandleWebSocket))
}