package models

import "gorm.io/gorm"

type GormMessage struct {
	gorm.Model
	Message  string `json:message`
	Reversed string `json:reversed`
}
