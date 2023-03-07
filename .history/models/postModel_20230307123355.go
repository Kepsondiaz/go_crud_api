package models 

import "gorm.io/gorm"

func pots() {
	type Posts struct {
		gorm.Model
		Title string
		Content  string
	  }
}