package models 

import "gorm.io/gorm"

func posts() {
	type User struct {
		gorm.Model
		Title string
		C string
	  }
}