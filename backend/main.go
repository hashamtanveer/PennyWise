package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"javascript.isdumb/pennywise/pkg/shared"
	"javascript.isdumb/pennywise/pkg/user"
)

func init() {
	godotenv.Load(".env")

	dbFilePath := ".db"
	if os.Getenv("DB_FILE_PATH") != "" {
		dbFilePath = os.Getenv("DB_FILE_PATH")
	}
	var err error
	shared.DB, err = gorm.Open(sqlite.Open(dbFilePath))
	if err != nil {
		panic("Could not open database file\n")
	}
	shared.DB.AutoMigrate(&user.User{}, &user.Transaction{})

	if os.Getenv("JWT_SECRET") != "" {
		shared.JwtSecret = []byte(os.Getenv("JWT_SECRET"))
	}
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
	}))

	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", user.UserLogin)
		userGroup.POST("/signup", user.UserSignup)
		//userGroup.GET("", user.AuthorizeMiddleware)
		//userGroup.PATCH("", user.AuthorizeMiddleware)
	}

	transactionsGroup := r.Group("/transactions")
	{
		transactionsGroup.GET("", user.AuthorizeMiddleware, user.GetTransactions)
		transactionsGroup.POST("", user.AuthorizeMiddleware, user.PostTransaction)
		transactionsGroup.GET("/:id", user.AuthorizeMiddleware)
		transactionsGroup.PATCH("/:id", user.AuthorizeMiddleware, user.PatchTransaction)
		transactionsGroup.DELETE("/:id", user.AuthorizeMiddleware, user.DeleteTransaction)
	}

	r.Run(":6900")
}
