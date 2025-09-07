package models

import "gorm.io/gorm"

type Funding struct {
	gorm.Model
	ProjectID uint    `json:"project_id"`
	UserID    uint    `json:"user_id"`
	Amount    float64 `json:"amount"`
}
