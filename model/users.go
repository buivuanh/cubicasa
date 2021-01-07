package model

import (
	"github.com/pkg/errors"
	"time"
)

type RoleType string

const AdminRole RoleType = "admin"
const ManagerRole RoleType = "manager"
const UserRole RoleType = "user"

type User struct {
	UserID    int       `json:"userId" db:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at,omitempty"`
	TeamID    int       `json:"teamId" db:"team_id"`
	Role      RoleType  `json:"role" db:"role"`
	Email     string    `json:"email" db:"email"`
}

func (u User) IsValid() error {
	if len(u.Email) == 0 {
		return errors.New("invalid user, email")
	}

	if len(u.Role) == 0 {
		return errors.New("invalid user, role")
	}

	if u.Role != AdminRole && u.Role != ManagerRole && u.Role != UserRole {
		return errors.New("invalid user, role")
	}

	return nil
}
