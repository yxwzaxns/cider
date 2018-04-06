package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yxwzaxns/cider/utils"
)

type Token struct {
	Token      string
	ExpireTime time.Time
}

var Tokens map[string]*Token

func Init() {
	Tokens = make(map[string]*Token)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// exclude auth url
		if c.Request.RequestURI == "/api/v1/user/auth" || c.Request.RequestURI == "/api/v1/ping" {
			c.Next()
			return
		}
		//
		println("token: ", c.Request.Header.Get("Authorization"))
		if c.Request.Header.Get("Authorization") != "" {
			if permit := CheckPermit(c.Request.Header.Get("Authorization")); !permit {
				c.AbortWithStatusJSON(200, gin.H{
					"code":   403,
					"status": "no access permit",
				})
			} else {
				UpdateTokenExpireTime()
				c.Next()
			}
		} else {
			c.AbortWithStatusJSON(200, gin.H{
				"code":   403,
				"status": "No Token fond!",
			})
		}
	}
}

func UpdateTokenExpireTime() {
	if utils.IsProduction() {
		Tokens["admin"].ExpireTime = time.Now().Local().Add(time.Duration(utils.GetTokenTimeOut() * time.Minute))
	}
}

func CheckPermit(token string) bool {
	if !utils.IsProduction() {
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
	Tokens["admin"] = &Token{Token: token, ExpireTime: time.Now().Local().Add(utils.GetTokenTimeOut() * time.Minute)}
	return token
}
