package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title string
	Url   string
	Info  string
}
