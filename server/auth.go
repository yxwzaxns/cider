package server

import (
	"cider/utils"
	"time"
)

type Token struct {
	Token      string
	ExpireTime time.Time
}

const TOKEN_TIMEOUT = 5

var Tokens = make(map[string]*Token)

func CheckPermit(token string) bool {
	if len(token) == 32 && len(Tokens) != 0 {
		if Tokens["admin"].Token == token && Tokens["admin"].ExpireTime.After(time.Now().Local()) {
			return true
		}
	}

	return false
}

func NewToken() string {
	token := utils.UUID()
	Tokens["admin"] = &Token{Token: token, ExpireTime: time.Now().Local().Add(TOKEN_TIMEOUT * time.Minute)}
	return token
}

func UpdateTokenExpireTime() {
	Tokens["admin"].ExpireTime = time.Now().Local().Add(TOKEN_TIMEOUT * time.Minute)
}
