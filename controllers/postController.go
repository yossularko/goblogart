package controllers

import (
	"goblogart/inits"
	"goblogart/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {

	var body models.Post
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Draft: body.Draft, Author: body.Author}

	err := inits.DB.Create(&post).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": post})

}

func GetPosts(ctx *gin.Context) {

	var posts []models.Post

	err := inits.DB.Find(&posts).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": posts})

}

func GetPost(ctx *gin.Context) {

	var post models.Post

	err := inits.DB.Where("id = ?", ctx.Param("id")).First(&post).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": post})

}

func UpdatePost(ctx *gin.Context) {

	var post models.Post

	err := inits.DB.Where("id = ?", ctx.Param("id")).First(&post).Error
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": post})

}

func DeletePost(ctx *gin.Context) {

	if err := inits.DB.Where("id = ?", ctx.Param("id")).Delete(&models.Post{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "post has been deleted successfully"})

}
