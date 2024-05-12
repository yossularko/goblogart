package controllers

import (
	"fmt"
	"goblogart/inits"
	"goblogart/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {

	var body models.Post
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	body.UserID = user.(models.User).ID
	body.Author = user.(models.User).Name

	post := models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Draft: body.Draft, Author: body.Author, UserID: body.UserID}

	err := inits.DB.Create(&post).Error
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": post})

}

func GetPosts(ctx *gin.Context) {

	var posts []models.Post

	search := fmt.Sprintf("%%%s%%", strings.ToLower(ctx.Query("search")))

	err := inits.DB.Find(&posts, "LOWER(title) LIKE ? OR LOWER(body) LIKE ?", search, search).Error
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": posts})

}

func GetPost(ctx *gin.Context) {

	var post models.Post

	err := inits.DB.Where("id = ?", ctx.Param("id")).First(&post).Error
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": post})

}

func GetMyPosts(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	UserID := user.(models.User).ID

	var posts []models.Post

	err := inits.DB.Where(&models.Post{UserID: UserID}).Find(&posts).Error
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": posts})

}

func UpdatePost(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	UserID := user.(models.User).ID

	var post models.Post

	err := inits.DB.Where("id = ?", ctx.Param("id")).Where(&models.Post{UserID: UserID}).First(&post).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body models.Post
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData := models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Draft: body.Draft, Author: body.Author}
	if err := inits.DB.Model(&post).Updates(updateData).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": post})

}

func DeletePost(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	UserID := user.(models.User).ID

	var post models.Post

	err := inits.DB.Where("id = ?", ctx.Param("id")).Where(&models.Post{UserID: UserID}).First(&post).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := inits.DB.Where("id = ?", ctx.Param("id")).Delete(&models.Post{}).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "post has been deleted successfully"})

}
