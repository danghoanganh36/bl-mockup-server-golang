package models

import (
	"time"
	"gorm.io/gorm"
)

type Author struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type Blog struct {
	gorm.Model
	CustomID   string    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `gorm:"embedded" json:"author"`
	Date       time.Time `json:"date"`
	Views      int       `json:"views"`
	Tags       []string  `gorm:"-" json:"tags"` 
	Excerpt    string    `json:"excerpt"`      
	Content    string    `json:"content"`
	Image      string    `json:"image"`
	URL        string    `json:"url"`
}	

type Category struct {
	gorm.Model
	Name  string
}

type Metric struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
	Blogs       []Blog `gorm:"many2many:metric_blogs;"`
}