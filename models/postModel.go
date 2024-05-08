package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string `json:"title"`
	Body   string `json:"body"`
	Likes  int    `json:"likes" gorm:"int;default:0"`
	Draft  bool   `json:"draft" gorm:"bool;default:false"`
	Author string `json:"author"`
}

type PostInput struct {
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	Likes  int    `json:"likes"`
	Draft  bool   `json:"draft"`
	Author string `json:"author" binding:"required"`
}
