package types

import "time"

type SignUpUser struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username,omitempty" binding:"omitempty,alphanum,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserBugsProject struct {
	ID           uint   `json:"id"`
	ProjectTitle string `json:"title"`
}

type UserBugsResponse struct {
	ID        uint            `json:"id"`
	Title     string          `json:"title"`
	Status    BugStatus       `json:"status"`
	Priority  Priority        `json:"priority"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Project   UserBugsProject `json:"project" gorm:"-"`
}

type UpdateUser struct {
	Name     *string `json:"name" binding:"omitempty"`
	Username *string `json:"username,omitempty" binding:"omitempty,alphanum,min=3,max=20"`
	Password *string `json:"password" binding:"omitempty,min=8"`
}
