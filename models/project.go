package models

import (
	"time"

	"gorm.io/gorm"
)

// Project represents a crowdfunding project. Status can be "draft" or "published".
type Project struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Goal        float64    `json:"goal"`
	Raised      float64    `json:"raised"`
	Deadline    string     `json:"deadline"`
	UserID      uint       `json:"user_id"`
	Status      string     `json:"status" gorm:"default:'draft'"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}
