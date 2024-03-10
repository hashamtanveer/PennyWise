package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    ID int `json:"id" gorm:"->;primaryKey;column:id"`
    CreatedAt *time.Time `json:"-" gorm:"autoCreateTime"`
    UpdatedAt *time.Time `json:"-" gorm:"autoUpdateTime"`
    DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`

    Username string `json:"username" gorm:"column:username;type:VARCHAR(64);uniqueIndex"`
    Password string `json:"-" gorm:"column:password;type:VARCHAR(72)"`
    Transactions []Transaction `json:"-" gorm:"foreignKey:UserID"`
}

type Transaction struct {
    ID int `json:"id" gorm:"->;primaryKey;column:id"`
    CreatedAt *time.Time `json:"-" gorm:"autoCreateTime"`
    UpdatedAt *time.Time `json:"-" gorm:"autoUpdateTime"`
    DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`

    UserID int `json:"user_id" gorm:"index;column:user_id"`
    Date *time.Time `json:"date" gorm:"column:date"`
    Description string `json:"description" gorm:"column:description;type:TEXT"`
    Category string `json:"category" gorm:"column:category;type:VARCHAR(64)"`
    Amount int `json:"amount" gorm:"column:amount;type:INTEGER"`
}

