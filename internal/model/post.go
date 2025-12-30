package model

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Content   string    `gorm:"type:text;not null"`
	AuthorID  uint      `gorm:"not null;index"`
	Status    string    `gorm:"type:varchar(20);default:'draft'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Post) TableName() string {
	return "posts"
}
