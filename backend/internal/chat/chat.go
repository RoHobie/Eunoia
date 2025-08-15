package chat

import (
	"encoding/json"
	"time"

	"eunoia/internal/models"
	"eunoia/internal/store"
	"eunoia/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func CreateRoom(c *fiber.Ctx) error {
	roomID := utils.GenerateRoomCode()
	err := store.Client.SAdd(store.Ctx, "room:"+roomID+":users", "").Err()
	if err != nil {
		return c.Status(500).SendString("Failed to create room")
	}
	store.Client.Expire(store.Ctx, "room:"+roomID+":users", time.Duration(store.GetRoomTTL())*time.Second)
	return c.JSON(fiber.Map{"roomID": roomID})
}

func JoinRoom(c *fiber.Ctx) error {
	var body struct {
		RoomID   string `json:"roomID"`
		Username string `json:"username"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	if exists, _ := store.Client.Exists(store.Ctx, "room:"+body.RoomID+":users").Result(); exists == 0 {
		return c.Status(404).SendString("Room not found")
	}

	store.Client.SAdd(store.Ctx, "room:"+body.RoomID+":users", body.Username)
	store.Client.Expire(store.Ctx, "room:"+body.RoomID+":users", time.Duration(store.GetRoomTTL())*time.Second)

	return c.JSON(fiber.Map{"joined": true})
}

func HandleWebSocket(conn *websocket.Conn) {
	roomID := conn.Params("roomID")
	username := conn.Query("username")

	// Add user to room
	store.Client.SAdd(store.Ctx, "room:"+roomID+":users", username)
	store.Client.Expire(store.Ctx, "room:"+roomID+":users", time.Duration(store.GetRoomTTL())*time.Second)

	defer func() {
		// Remove user
		store.Client.SRem(store.Ctx, "room:"+roomID+":users", username)

		// If room is empty, delete all keys
		users, _ := store.Client.SMembers(store.Ctx, "room:"+roomID+":users").Result()
		if len(users) == 0 {
			store.Client.Del(store.Ctx, "room:"+roomID+":users")
			store.Client.Del(store.Ctx, "room:"+roomID+":messages")
		}
		conn.Close()
	}()

	// Send last 50 messages
	msgs, _ := store.Client.LRange(store.Ctx, "room:"+roomID+":messages", -50, -1).Result()
	for _, m := range msgs {
		conn.WriteMessage(websocket.TextMessage, []byte(m))
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		message := models.Message{
			Username:  username,
			Content:   string(msg),
			Timestamp: time.Now().Format(time.RFC3339),
		}

		jsonMsg, _ := json.Marshal(message)

		// Store in Redis list (last 50 messages)
		store.Client.RPush(store.Ctx, "room:"+roomID+":messages", jsonMsg)
		store.Client.LTrim(store.Ctx, "room:"+roomID+":messages", -50, -1)
		store.Client.Expire(store.Ctx, "room:"+roomID+":messages", time.Duration(store.GetRoomTTL())*time.Second)

		// Broadcast to all users in room
		// (Here we would have a proper connection registry if we wanted fanout without polling)
		conn.WriteMessage(websocket.TextMessage, jsonMsg)
	}
}
