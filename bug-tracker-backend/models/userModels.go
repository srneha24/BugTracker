package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `json:"username" gorm:"unique;not null"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"password" gorm:"not null"`
	Projects []Project `json:"projects" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Projects created by the user
	Teams    []Team    `json:"teams" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`        // Teams the user is part of
	Bugs     []Bug     `json:"bugs" gorm:"foreignKey:AssignedTo;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`    // Bugs assigned to the user
}
