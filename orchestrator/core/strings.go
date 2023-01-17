package core

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/google/uuid"
)

func HashString(input string) string {
	bytes := sha1.Sum([]byte(input))

	return hex.EncodeToString(bytes[:])
}

func GenerateUUID() string {
	id := uuid.New()

	return id.String()
}
