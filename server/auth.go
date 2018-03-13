package server

import (
	G "cider/global"
	"cider/utils"
	"time"
)

type Token struct {
	Token      string
	ExpireTime time.Time
}

const DEBUG = true

var Tokens = make(map[string]*Token)

func CheckPermit(token string) bool {
	if DEBUG == true {
		return true
	}
	if len(token) == 32 && len(Tokens) != 0 {
		if Tokens["admin"].Token == token && Tokens["admin"].ExpireTime.After(time.Now().Local()) {
			return true
		}
	}

	return false
}

func NewToken() string {
	token := utils.UUID()
	Tokens["admin"] = &Token{Token: token, ExpireTime: time.Now().Local().Add(time.Duration(G.Config.TokenTimeout) * time.Minute)}
	return token
}

func UpdateTokenExpireTime() {
	if DEBUG != true {
		Tokens["admin"].ExpireTime = time.Now().Local().Add(time.Duration(G.Config.TokenTimeout) * time.Minute)
	}

}
