package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

type SHA256Hasher struct{}

func (h *SHA256Hasher) HashPassword(password string) (string, error) {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

func (h *SHA256Hasher) CheckPassword(password, hash string) bool {
	hashedPassword, err := h.HashPassword(password)
	if err != nil {
		return false
	}
	return hashedPassword == hash
}
