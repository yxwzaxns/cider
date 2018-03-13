package utils

import (
	"encoding/base64"
	"encoding/hex"
	"os"
	"strconv"

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

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
