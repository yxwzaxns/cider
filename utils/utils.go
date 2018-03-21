package utils

import (
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

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

func OpenFile(path string) io.Writer {
	if f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		log.Fatal(err)
	} else {
		return f
	}
	return nil
}

func ParseField(field string) (k string, v interface{}) {
	kv := strings.Split(field, ":")
	switch kv[1] {
	case "1":
		return kv[0], true
	case "0":
		return kv[0], false
	default:
		return kv[0], kv[1]
	}
}
