package models 

import "gorm.io/gorm"

func posts() {
	type Posts struct {
		gorm.Model
		Title string
		Content  string
	  }
}