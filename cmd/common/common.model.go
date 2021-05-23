package common

import "time"

type Model struct {
	ID        uint       `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"primary_key;column:id" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"primary_key;column:id" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"primary_key;column:id" json:"deletedAt"`
}
