package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    ID uint `json:"id" gorm:"->;primaryKey;column:id"`
    CreatedAt *time.Time `json:"-" gorm:"autoCreateTime"`
    UpdatedAt *time.Time `json:"-" gorm:"autoUpdateTime"`
    DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`

    Username string `json:"username" gorm:"column:username;type:VARCHAR(64);uniqueIndex"`
    Password string `json:"-" gorm:"column:password;type:VARCHAR(72)"`
}
