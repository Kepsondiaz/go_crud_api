package models 

import "gorm.io/gorm"

func posts() {
	type  struct {
		gorm.Model
		Title string
		Content  string
	  }
}