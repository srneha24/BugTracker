package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name" gorm:"not null;type:varchar(100)"`
	Username string    `json:"username" gorm:"unique;not null;type:varchar(100)"`
	Email    string    `json:"email" gorm:"unique;not null;type:varchar(100)"`
	Password string    `json:"password" gorm:"not null"`
	Projects []Project `json:"projects" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Projects created by the user
	Teams    []Team    `json:"teams" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`        // Teams the user is part of
	Bugs     []Bug     `json:"bugs" gorm:"foreignKey:AssignedTo;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`    // Bugs assigned to the user
}
