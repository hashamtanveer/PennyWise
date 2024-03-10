package user

import (
	"time"

	"gorm.io/gorm"

    "javascript.isdumb/pennywise/pkg/transactions"
)

type User struct {
    ID int `json:"id" gorm:"->;primaryKey;column:id"`
    CreatedAt *time.Time `json:"-" gorm:"autoCreateTime"`
    UpdatedAt *time.Time `json:"-" gorm:"autoUpdateTime"`
    DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`

    Username string `json:"username" gorm:"column:username;type:VARCHAR(64);uniqueIndex"`
    Password string `json:"-" gorm:"column:password;type:VARCHAR(72)"`
    Transactions []transactions.Transaction `json:"-" gorm:"foreignKey:UserID"`
}
