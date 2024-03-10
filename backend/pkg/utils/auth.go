package utils

import "github.com/gin-gonic/gin"

func GetUserIDFromContext(c *gin.Context) (uint64, bool) {
    userID, exists := c.Get("authorized_user_id")
    userIDInt, ok := userID.(uint64)
    return userIDInt, exists && ok
}
