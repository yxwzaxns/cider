package utils

import (
	"encoding/base64"
	"encoding/hex"
	"os"

	"github.com/satori/go.uuid"
)

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}

func GetKey() string {
	return os.Getenv("CIDER_AUTH_KEY")
}

func UUID() string {
	return hex.EncodeToString(uuid.Must(uuid.NewV4()).Bytes())
}
