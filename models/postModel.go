package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	Likes  int    `json:"likes" gorm:"int;default:0"`
	Draft  bool   `json:"draft" gorm:"bool;default:true"`
	Author string `json:"author" binding:"required"`
}
