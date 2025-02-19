package model

import (
	"time"

	"github.com/google/uuid"
)

type Listing struct {
	ID          uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`       // Unique ID for each listing
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`  // Title of the listing
	Description string    `json:"description" gorm:"type:text"`             // Details about the land
	Location    string    `json:"location" gorm:"type:varchar(255);not null"` // Physical location
	Size        float64   `json:"size" gorm:"type:decimal(10,2);not null"`  // Size in acres or hectares
	Price       float64   `json:"price" gorm:"type:decimal(12,2);not null"` // Price of the land
	ImageURL    string    `json:"image_url" gorm:"type:varchar(500)"`       // URL of the land image
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`         // Timestamp for creation
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`         // Timestamp for updates
}


