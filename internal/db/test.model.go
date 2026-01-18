package db

import "time"

type TestModel struct {
	ID        uint `gorm:"primaryKey"`
	Message   string
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

