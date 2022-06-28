package models

// import (
// 	""
// )

type Task struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Description string `gorm:"not null" json:"description" form:"description" valid:"required~Description is required"`
	Status      bool   `gorm:"not null" json:"status" form:"status" valid:"required~Status is required"`
	UserID      int
	User        *User
	CategoryID  int
	Category    *Category
	TimeModel
}
