package utils

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateRoomCode() string {
	return strings.Split(uuid.New().String(),"-")[0]
}