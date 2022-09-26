package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	Name      string         `json:"name"`
	Country   string         `json:"country"`
	Website   string         `json:"website"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
