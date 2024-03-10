package transactions

import (
    "time"
    "gorm.io/gorm"
)

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
