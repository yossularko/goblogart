package middlewares

import (
	"errors"
	"fmt"
	"goblogart/inits"
	"goblogart/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	tokenString, errCookie := ctx.Cookie("Authorization")

	if errCookie != nil {
		hAuth := ctx.Request.Header.Get("Authorization")

		splitedBearer := strings.Split(hAuth, "Bearer ")
		if len(splitedBearer) < 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString = splitedBearer[1]

	}

	token, errParse := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errVal := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New(errVal)

		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if errParse != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token parse"})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		if errGetUsr := inits.DB.First(&user, int(claims["id"].(float64))).Error; errGetUsr != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user)
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.Next()
}
