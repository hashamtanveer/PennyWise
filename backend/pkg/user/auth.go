package user

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
	"gorm.io/gorm"

	"javascript.isdumb/pennywise/pkg/shared"
	"javascript.isdumb/pennywise/pkg/utils"
)

var ErrUserExisted error = errors.New("UserExisted")
var ErrUserNotExisted error = errors.New("UserNotExisted")
var ErrWrongPassword error = errors.New("WrongPassword")
var ErrBadToken error = errors.New("BadToken")

type authRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type authResponse struct {
	ErrorStatus  int    `json:"error_status"`
	ErrorMessage string `json:"error_message"`
	Token        string `json:"token"`
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

func genToken(userID int) (string, error) {
	timeNow := time.Now()
	timeExpire := timeNow.Add(time.Hour * 24)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userID),
		Issuer:    "PennyWise",
		IssuedAt:  jwt.NewNumericDate(timeNow),
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

	token, err := genToken(user.ID)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, authResponse{0, "", token})
}

func AuthorizeMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	if len(tokenString) <= 0 {
		utils.ResponseWithError(c, http.StatusUnauthorized, ErrBadToken)
		c.Abort()
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(shared.JwtSecret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		utils.ResponseWithError(c, http.StatusUnauthorized, err)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		utils.ResponseWithError(c, http.StatusUnauthorized, ErrBadToken)
		c.Abort()
		return
	}

	if claims.NotBefore.After(time.Now()) || time.Now().After(claims.ExpiresAt.Time) {
		utils.ResponseWithError(c, http.StatusUnauthorized, ErrBadToken)
		c.Abort()
		return
	}

	userID, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		utils.ResponseWithError(c, http.StatusUnauthorized, ErrBadToken)
		c.Abort()
		return
	}
	c.Set("authorized_user_id", userID)
	c.Next()
}
