// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import (
	"time"
)

// User -.
type User struct {
	ID        string    `json:"id" gorm:"column:id;unique;primaryKey"`
	Name      string    `json:"name" gorm:"column:nickName"`
	Age       string    `json:"age"  gorm:"column:userType"`
	CreatedAt time.Time `json:"created_at" gorm:"column:createDate"`
}

func (User) TableName() string {
	return "userProfile"
}
