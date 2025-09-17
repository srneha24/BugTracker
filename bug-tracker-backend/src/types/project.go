package types

import (
	"time"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/utils"
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
	TeamRoleAdmin     TeamRole = "admin"
	TeamRoleDeveloper TeamRole = "dev"
	TeamRoleTester    TeamRole = "tester"
)

func (t TeamRole) Value() string {
	return string(t)
}

type Priority uint

const (
	PriorityHigh Priority = iota + 1
	PriorityMedium
	PriorityLow
)

func (p Priority) Value() uint {
	return uint(p)
}

type BugStatus string

const (
	BugStatusTodo       BugStatus = "todo"
	BugStatusInProgress BugStatus = "in_progress"
	BugStatusDone       BugStatus = "done"
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
	Priority    Priority  `json:"priority" binding:"omitempty,oneof=1 2 3"` // 1: High, 2: Medium, 3: Low
	AssignedTo  uint      `json:"assigned_to" binding:"required"`
}

// SetDefaults sets default values for CreateBug fields if they are omitted.
func (c *CreateBug) SetDefaults() {
	if c.Status == "" {
		c.Status = BugStatusTodo
	}

	if c.Priority == 0 {
		c.Priority = PriorityHigh
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

type UpdateProject struct {
	Title       *string `json:"title" binding:"omitempty"`
	Description *string `json:"description" binding:"omitempty"`
}

type BugListQueryParams struct {
	*utils.PaginationQueryParams
	Search     *string `form:"search"`
	Tags       *string `form:"tags"`
	Deadline   *string `form:"deadline"`
	Status     *string `form:"status" binding:"omitempty,oneof=todo in_progress done"`
	Priority   *int    `form:"priority" binding:"omitempty,oneof=1 2 3"`
	AssignedTo *uint   `form:"assigned_to"`
}
