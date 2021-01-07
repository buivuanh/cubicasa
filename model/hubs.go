package model

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

type GeoLocation struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func (m GeoLocation) Value() (driver.Value, error) {
	res, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return string(res), nil
}

func (m *GeoLocation) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	v, ok := value.([]byte)
	if !ok {
		return errors.New("type of value is not bytes")
	}

	err := json.Unmarshal(v, m)
	if err != nil {
		return err
	}

	return nil
}

type Hub struct {
	ID          int         `json:"id" db:"id,omitempty"`
	Name        string      `json:"name" db:"name"`
	CreatedAt   time.Time   `json:"createdAt" db:"created_at,omitempty"`
	UpdatedAt   time.Time   `json:"updatedAt" db:"updated_at,omitempty"`
	GeoLocation GeoLocation `json:"geoLocation" db:"geo_location"`
}

func (h Hub) IsValid() error {
	if len(h.Name) == 0 {
		return errors.New("invalid Hub, name")
	}

	return nil
}

