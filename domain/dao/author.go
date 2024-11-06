package dao

import (
	"base-gin/domain"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID        uint 				 `gorm:"primarykey"`
	Fullname  string             `gorm:"size:56;not null;"`
	Gender    *domain.TypeGender `gorm:"type:enum('f','m');"`
	BirthDate *time.Time
}

func (Author) TableName() string {
	return "authors"
}
