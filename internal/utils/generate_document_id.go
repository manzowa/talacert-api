package utils

import (
	"fmt"
	"strings"
)

var documentPrefixes = map[string]string{
	"certificat": "CERT",
	"diplôme":    "DIP",
}

func GenerateDocumentID(documentType string, year int, seq int) string {
	prefix, ok := documentPrefixes[strings.ToLower(documentType)]

	if !ok {
		// Par défaut si le type de document n'est pas reconnu
		prefix = "DOC"
	}
	return fmt.Sprintf("%s-%d-%04d", prefix, year, seq)
}
