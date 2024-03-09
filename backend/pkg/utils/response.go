package utils

import (

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
    ErrorMessage string `json:"error_message"`
}

func ResponseWithError(c *gin.Context, status int, err error) {
    c.JSON(status, BaseResponse{
        err.Error(),
    })
}
