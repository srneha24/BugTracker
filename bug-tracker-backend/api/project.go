package api

import "github.com/WNBARookie/BugTracker/bug-tracker-backend/utils"

type CreateProject struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
}

type ProjectResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedBy   int    `json:"created_by"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
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
