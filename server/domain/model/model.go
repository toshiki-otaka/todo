package model

import "time"

type Item struct {
	Id         int       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title      string    `gorm:"not null" json:"title"`
	Desc       string    `json:"desc"`
	IsDone     bool      `gorm:"not null;default false" json:"is_done"`
	InsertedAt time.Time `gorm:"not null" json:"inserted_at"`
	UpdatedAt  time.Time `gorm:"not null" json:"updated_at"`
}
