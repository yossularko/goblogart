package controllers

import (
	"goblogart/inits"
	"goblogart/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {

	var body models.User

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}

	errCreate := inits.DB.Create(&user).Error
	if errCreate != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": errCreate.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(ctx *gin.Context) {

	var body struct {
		Email    string `json:"email" binding:"required" gorm:"unique"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := inits.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": user.ID, "exp": time.Now().Add(time.Hour * 24 * 30).Unix()})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error signing token"})
		return
	}

	newUser := user
	newUser.Password = ""

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"data": "success login", "user": newUser})
}

func GetUsers(ctx *gin.Context) {
	var users []models.User

	err := inits.DB.Find(&users).Error
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var newUser []models.User

	for _, usr := range users {
		usr.Password = ""
		newUser = append(newUser, usr)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newUser})
}

func Validate(ctx *gin.Context) {
	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user not exist"})
		return
	}

	newUser := user.(models.User)
	newUser.Password = ""

	ctx.JSON(http.StatusOK, gin.H{"data": "You are logged in!", "user": newUser})
}

func Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"data": "You are logged out!"})
}
