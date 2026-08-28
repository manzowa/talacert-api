package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"io"
	"os"
)

type HashGenerator struct{}

var documentPrefixes = map[string]string{
	"certificat": "CERT",
	"diplôme":    "DIP",
}

func NewHashGenerator() *HashGenerator {
	return &HashGenerator{}
}

// GenerateHash creates a SHA256 hash from a string
func (h *HashGenerator) Generate(value string) string {

	hash := sha256.Sum256([]byte(value))

	return hex.EncodeToString(hash[:])
}

func (h *HashGenerator) GenerateFile(path string) (string, error) {

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

func (h *HashGenerator) Check(content, hash string) bool {

	generatedHash := h.Generate(content)

	return generatedHash == hash
}

func (h *HashGenerator) GenerateDocumentID(documentType string, year int, seq int) string {
	prefix, ok := documentPrefixes[strings.ToLower(documentType)]

	if !ok {
		// Par défaut si le type de document n'est pas reconnu
		prefix = "DOC"
	}
	return fmt.Sprintf("%s-%d-%04d", prefix, year, seq)
}
