package models

// import (
// 	""
// )

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Fullname string `gorm:"not null;uniqueIndex" json:"full_name" form:"full_name" valid:"required~Title is required"`
	Email     string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Email is invalid"`
	Password  string `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to have minimum length 6 characters"`
	Role string	`gorm:"not null" json:"role" form:"role" valid:"required~Role is requiered"`
	TimeModel
}
