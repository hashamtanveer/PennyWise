package transactions

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"javascript.isdumb/pennywise/pkg/shared"
	"javascript.isdumb/pennywise/pkg/utils"
)

func getTransactions(c *gin.Context) {
	userID, exists := utils.GetUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error_message": "InternalError"})
		return
	}

	var transactions []Transaction
	if result := shared.DB.
		Model(&Transaction{}).
		Where("user_id = ?", userID).
		Find(&transactions); result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
            utils.ResponseWithError(c, http.StatusInternalServerError, result.Error)
        }

        return
	}

    c.JSON(http.StatusOK, transactions)
}

func postTransaction(c *gin.Context) {

}
