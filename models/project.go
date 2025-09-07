package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Goal        float64 `json:"goal"`
	Raised      float64 `json:"raised"`
	Deadline    string  `json:"deadline"`
	UserID      uint    `json:"user_id"`
}
