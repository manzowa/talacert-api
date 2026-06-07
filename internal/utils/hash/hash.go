package hash

import (
	"crypto/sha256"
	"encoding/hex"

	"io"
	"os"
)

type Hash struct{}

func NewHash() *Hash {
	return &Hash{}
}

// GenerateHash creates a SHA256 hash from a string
func (h *Hash) GenerateHash(content string) string {

	hash := sha256.Sum256([]byte(content))

	return hex.EncodeToString(hash[:])
}

func (h *Hash) GenerateFileHash(path string) (string, error) {

	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hasher := sha256.New()

	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func (h *Hash) VerifyHash(content, hash string) bool {

	generatedHash := h.GenerateHash(content)

	return generatedHash == hash
}
