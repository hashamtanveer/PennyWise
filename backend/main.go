package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
    "github.com/glebarez/sqlite"

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

    if os.Getenv("JWT_SECRET") != "" {
        shared.JwtSecret = []byte(os.Getenv("JWT_SECRET"))
    }
}

func main() {
    r := gin.Default()
    userGroup := r.Group("/user")
    {
        userGroup.POST("/login", user.UserLogin)
        userGroup.POST("/signup", user.UserSignup)
        userGroup.GET("")
        userGroup.PATCH("")
    }

    r.Run(":6969")
}
