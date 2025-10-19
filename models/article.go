package models

import "time"

type Article struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title    string `json:"title" gorm:"type:varchar(200);not null"`
	Content  string `json:"content" gorm:"type:text;not null"`
	Category string `json:"category" gorm:"type:varchar(100);not null"`
	Status   string `json:"status" gorm:"type:varchar(50);not null"`

	CreatedDate time.Time `json:"created_date" gorm:"autoCreateTime"`
	UpdatedDate time.Time `json:"updated_date" gorm:"autoUpdateTime"`
}
