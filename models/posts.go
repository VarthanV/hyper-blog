package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"unique" json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}
