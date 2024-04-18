package middleware

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
	"webook/webook/internal/web"
)

type LoginMiddlewareJWTBuilder struct {
}

func NewLoginMiddlewareJWTBuilder() *LoginMiddlewareJWTBuilder {
	return &LoginMiddlewareJWTBuilder{}
}

func (l *LoginMiddlewareJWTBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		// no need to login and check
		if ctx.Request.URL.Path == "/users/login" ||
			ctx.Request.URL.Path == "/users/signup" {
			return
		}

		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segments := strings.Split(tokenHeader, " ")
		if len(segments) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segments[1]
		userClaim := &web.UserClaim{}
		token, err := jwt.ParseWithClaims(tokenStr, userClaim, func(token *jwt.Token) (interface{}, error) {
			return web.JWTKey, nil
		})
		if err != nil || !token.Valid || userClaim.UserId == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if userClaim.UserAgent != ctx.Request.UserAgent() {
			// security problem, need to monitor it
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		now := time.Now()
		if userClaim.ExpiresAt.Sub(now) < 20*time.Minute {
			userClaim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour))
			tokenStr, err = token.SignedString(web.JWTKey)
			ctx.Header("x-jwt-token", tokenStr)
			if err != nil {
				// 这边不要中断，因为仅仅是过期时间没有刷新，但是用户是登录了的
				log.Println(err)
			}
		}
		ctx.Set("userClaim", userClaim)
	}
}
