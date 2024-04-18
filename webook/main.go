package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
	"webook/webook/config"
	"webook/webook/internal/repository"
	"webook/webook/internal/repository/dao"
	"webook/webook/internal/service"
	"webook/webook/internal/web"
	"webook/webook/internal/web/middleware"
)

func main() {
	db := initDB()
	server := initWebServer()

	u := initUser(db)
	u.RegisterRoutes(server)

	//server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	server.Run(":8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()

	//redisClient := redis.NewClient(&redis.Options{
	//	Addr: config.Config.Redis.Addr,
	//})
	//server.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())

	// CORS (Cross-Origin Resource Sharing)
	server.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"x-jwt-token"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return true // strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	//store := cookie.NewStore([]byte("secret"))
	//store, err := redis.NewStore(16, "tcp", "localhost:6379", "",
	//	// authentication key, encryption key
	//	[]byte("sqVQkXPKmOdc9p6Lj5n5JEvMEMkj16N8"),
	//	[]byte("clm53EGOXFDJLtTknh8bisc2jrKzHbNO"),
	//)
	//if err != nil {
	//	panic(err)
	//}

	//store := memstore.NewStore(
	//	[]byte("sqVQkXPKmOdc9p6Lj5n5JEvMEMkj16N8"),
	//	[]byte("clm53EGOXFDJLtTknh8bisc2jrKzHbNO"),
	//)

	//server.Use(sessions.Sessions("ssid", store))
	//server.Use(middleware.NewLoginMiddlewareBuilder().Build())
	server.Use(middleware.NewLoginMiddlewareJWTBuilder().Build())

	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	println(config.Config.DB.DSN)
	if err != nil {
		panic(err)
	}

	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
