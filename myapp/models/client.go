package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name string `json:"name"`
	Info string `json:"info"`
}