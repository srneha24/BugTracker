package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null;type:varchar(100)"`
	Description string `json:"description"`
	CreatedBy   uint   `json:"created_by" gorm:"not null"`
	User        User   `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // User who created the project
}

type Team struct {
	gorm.Model
	ProjectID uint    `json:"project_id" gorm:"not null;uniqueIndex:idx_project_user"`                                   // Unique index to prevent duplicate team members in the same project
	Project   Project `json:"-" gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Project associated with the team
	UserID    uint    `json:"user_id" gorm:"not null;uniqueIndex:idx_project_user"`                                      // Unique index to prevent duplicate team members in the same project
	User      User    `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`    // User associated with the team
	Role      string  `json:"role" gorm:"not null"`
}

type Bug struct {
	gorm.Model
	Title        string         `json:"title" gorm:"not null;type:varchar(100)"`
	Description  string         `json:"description"`
	Tags         pq.StringArray `json:"tags" gorm:"type:varchar[]"`
	Deadline     time.Time      `json:"deadline" gorm:"not null"`
	Status       string         `json:"status" gorm:"not null;default:'todo'"` // todo, in_progress, done
	Priority     uint           `json:"priority" gorm:"not null"`              // 1: High, 2: Medium, 3: Low
	AssignedTo   uint           `json:"assigned_to" gorm:"not null"`
	AssignedUser User           `json:"-" gorm:"foreignKey:AssignedTo;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // User to whom the bug is assigned
	ProjectID    uint           `json:"project_id" gorm:"not null"`
	Project      Project        `json:"-" gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Project associated with the bug
}
