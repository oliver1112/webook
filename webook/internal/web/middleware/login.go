package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		// no need to login and check
		if ctx.Request.URL.Path == "/users/login" ||
			ctx.Request.URL.Path == "/users/signup" {
			return
		}
		sess := sessions.Default(ctx)
		userID := sess.Get("userId")
		if userID == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		now := time.Now()

		// 我怎么知道，要刷新了呢？
		// 假如说，我们的策略是每分钟刷一次，我怎么知道，已经过了一分钟？
		const updateTimeKey = "update_time"
		// 试着拿出上一次刷新时间
		val := sess.Get(updateTimeKey)
		lastUpdateTime, ok := val.(time.Time)
		if val == nil || !ok || now.Sub(lastUpdateTime) > time.Minute*5 {
			// 你这是第一次进来
			sess.Set(updateTimeKey, now)
			sess.Options(sessions.Options{
				// second
				MaxAge: 10 * 60,
			})
			sess.Set("userId", userID)
			err := sess.Save()
			if err != nil {
				// 打日志
				fmt.Println(err)
			}
		}
	}
}
