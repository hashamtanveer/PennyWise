package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
	"gorm.io/gorm"
    "github.com/golang-jwt/jwt/v4"

	"javascript.isdumb/pennywise/pkg/shared"
	"javascript.isdumb/pennywise/pkg/utils"
)

var ErrUserExisted error = errors.New("UserExisted")
var ErrUserNotExisted error = errors.New("UserNotExisted")
var ErrWrongPassword error = errors.New("WrongPassword")

type authRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type authResponse struct {
    ErrorStatus int `json:"error_status"`
    ErrorMessage string `json:"error_message"`
    Token string `json:"token"`
}

func UserSignup(c *gin.Context) {
	requestBody := authRequest{}
	if err := c.ShouldBind(&requestBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, err)
		return
	}

	pwHashed, err := bcrypt.GenerateFromPassword([]byte(sha3.New256().Sum([]byte(requestBody.Password))), bcrypt.DefaultCost)
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, err)
		return
	}

	newUser := User{
		Username: requestBody.Username,
		Password: string(pwHashed),
	}
	if result := shared.DB.Model(&User{}).Create(&newUser); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			utils.ResponseWithError(c, http.StatusBadRequest, ErrUserExisted)
		} else {
			utils.ResponseWithError(c, http.StatusInternalServerError, result.Error)
		}

		return
	}

	c.String(http.StatusOK, "OK")
}

func genToken(username string) (string, error) {
    timeNow := time.Now()
    timeExpire := timeNow.Add(time.Hour * 24)

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
        Subject: username,
        Issuer: "PennyWise",
        IssuedAt: jwt.NewNumericDate(timeNow),
        ExpiresAt: jwt.NewNumericDate(timeExpire),
        NotBefore: jwt.NewNumericDate(timeNow),
    })

    return token.SignedString(shared.JwtSecret)
}

func UserLogin(c *gin.Context) {
	requestBody := authRequest{}
	if err := c.ShouldBind(&requestBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, err)
		return
	}

	user := User{}
	if result := shared.DB.
		Model(&User{}).
		Where("username = ?", requestBody.Username).
		First(&user); result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            utils.ResponseWithError(c, http.StatusUnauthorized, ErrUserNotExisted)
        } else {
            utils.ResponseWithError(c, http.StatusInternalServerError, result.Error)
        }

        return
	}

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), sha3.New256().Sum([]byte(requestBody.Password))); err != nil {
        utils.ResponseWithError(c, http.StatusUnauthorized, ErrWrongPassword)
        return
    }

    token, err := genToken(requestBody.Username)
    if err != nil {
        utils.ResponseWithError(c, http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, authResponse{0,"",token})
}
