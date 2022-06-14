package models

type Category struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Type string `gorm:"not null;uniqueIndex" json:"type" form:"type" valid:"required~Type is required"`
	TimeModel
}
