package api

import (
	"time"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/utils"
)

type CreateProject struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
}

type ProjectResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProjectListQueryParams struct {
	*utils.PaginationQueryParams
	Search *string `form:"search"`
}

type TeamRole string

const (
	Admin     TeamRole = "admin"
	Developer TeamRole = "dev"
	Tester    TeamRole = "tester"
)

func (t TeamRole) Value() string {
	return string(t)
}

type Priority uint

const (
	High Priority = iota + 1
	Medium
	Low
)

func (p Priority) Value() uint {
	return uint(p)
}

type BugStatus string

const (
	Todo       BugStatus = "tod"
	InProgress BugStatus = "in_progress"
	Done       BugStatus = "done"
)

func (s BugStatus) Value() string {
	return string(s)
}

type ProjectURI struct {
	ProjectID uint `uri:"projectID" binding:"required"`
}

type CreateBug struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"omitempty"`
	Tags        []string  `json:"tags" binding:"omitempty"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	Status      BugStatus `json:"status" binding:"omitempty,oneof=todo in_progress done"`
	Priority    Priority  `json:"priority" binding:"required"`
	AssignedTo  uint      `json:"assigned_to" binding:"required"`
}

// SetDefaults sets default values for CreateBug fields if they are omitted.
func (c *CreateBug) SetDefaults() {
	if c.Status == "" {
		c.Status = Todo
	}
}

type AssignedTo struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BugResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Tags        []string   `json:"tags"`
	Deadline    time.Time  `json:"deadline"`
	Status      BugStatus  `json:"status"`
	Priority    Priority   `json:"priority"`
	AssignedTo  AssignedTo `json:"assigned_to"`
	ProjectID   uint       `json:"project_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type BugURI struct {
	BugID uint `uri:"bugID" binding:"required"`
}
