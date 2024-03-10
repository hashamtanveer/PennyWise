package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"javascript.isdumb/pennywise/pkg/shared"
	"javascript.isdumb/pennywise/pkg/utils"
)

type transactionUpdate struct {
	ID        int             `json:"-"`
	CreatedAt *time.Time      `json:"-"`
	UpdatedAt *time.Time      `json:"-"`
	DeletedAt *gorm.DeletedAt `json:"-"`

	UserID      int        `json:"-"`
	Date        *time.Time `json:"date"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	Amount      int        `json:"amount"`
}

func GetTransactions(c *gin.Context) {
	userID, exists := utils.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error_message": "InternalError"})
		return
	}

	var user User
	if result := shared.DB.
		Model(&User{}).
		Preload("Transactions").
		Where("id = ?", userID).
		Find(&user); result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.ResponseWithError(c, http.StatusInternalServerError, result.Error)
		}

		return
	}

	c.JSON(http.StatusOK, user.Transactions)
}

func PostTransaction(c *gin.Context) {
	userID, exists := utils.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error_message": "InternalError"})
		return
	}

	newTransaction := transactionUpdate{}
	if err := c.ShouldBind(&newTransaction); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, err)
		return
	}

	stuff := Transaction(newTransaction)
	stuff.UserID = int(userID)
	if result := shared.DB.
		Create(&stuff); result.Error != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, result.Error)
		return
	}

	c.String(http.StatusOK, "OK")
}

func PatchTransaction(c *gin.Context) {
	userID, exists := utils.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error_message": "InternalError"})
		return
	}
	id := c.Param("id")

	updatedTransaction := transactionUpdate{}
	if err := c.ShouldBind(&updatedTransaction); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, err)
		return
	}

	stuff := Transaction(updatedTransaction)
	if result := shared.DB.
		Model(&Transaction{}).
		Where("user_id = ? AND id = ?", userID, id).
		Updates(&stuff); result.Error != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, result.Error)
		return
	}

	c.String(http.StatusOK, "OK")
}

func DeleteTransaction(c *gin.Context) {
	userID, exists := utils.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error_message": "InternalError"})
		return
	}
	id := c.Param("id")

    if result := shared.DB.
        Where("id = ? AND user_id = ?", id, userID).
        Delete(&Transaction{}); result.Error != nil {
            utils.ResponseWithError(c, http.StatusInternalServerError, result.Error)
            return
        }

    c.String(http.StatusOK, "OK")
}
