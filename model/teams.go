package model

import (
	"github.com/pkg/errors"
	"time"
)

type TeamType string

const (
	IT   TeamType = "it"
	SALE TeamType = "sale"
)

type ProductionTeam struct {
	ID        int       `json:"id" db:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at,omitempty"`
	Name      string    `json:"name" db:"name"`
	TeamType  TeamType  `json:"teamType" db:"team_type"`
	HubID     *int      `json:"hubId" db:"hub_id"`
}

func (pt ProductionTeam) IsValid() error {
	if len(pt.Name) == 0 {
		return errors.New("invalid ProductionTeam, name")
	}

	if len(pt.TeamType) == 0 {
		return errors.New("invalid ProductionTeam, name")
	}

	return nil
}
