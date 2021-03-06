package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Container is used by pop to map your containers database table to your go code.
type Container struct {
	ID              uuid.UUID    `json:"id" db:"id"`
	CreatedAt       time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at" db:"updated_at"`
	CreatedBy       uuid.UUID    `json:"created_by" db:"created_by"`
	TenantID        uuid.UUID    `json:"tenant_id" db:"tenant_id"`
	CarrierID       nulls.UUID   `json:"carrier_id" db:"carrier_id"`
	TerminalID      nulls.UUID   `json:"terminal_id" db:"terminal_id"`
	YardID          nulls.UUID   `json:"yard_id" db:"yard_id"`
	OrderID         nulls.UUID   `json:"order_id" db:"order_id"`
	SerialNumber    nulls.String `json:"serial_number" db:"serial_number"`
	Origin          nulls.String `json:"origin" db:"origin"`
	Destination     nulls.String `json:"destination" db:"destination"`
	Lfd             nulls.Time   `json:"lfd" db:"lfd"`
	ReservationTime nulls.Time   `json:"reservation_time" db:"reservation_time"`
	Size            nulls.String `json:"size" db:"size"`
	Type            nulls.String `json:"type" db:"type"`
	Status          nulls.String `json:"status" db:"status"`
	DriverID        nulls.UUID   `json:"driver_id" db:"driver_id"`
	GpsURL          nulls.String `json:"gps_url" db:"gps_url"`
}

// String is not required by pop and may be deleted
func (c Container) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Containers is not required by pop and may be deleted
type Containers []Container

// String is not required by pop and may be deleted
func (c Containers) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Container) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.SerialNumber.String, Name: "SerialNumber"},
		&validators.StringIsPresent{Field: c.Origin.String, Name: "Origin"},
		&validators.StringIsPresent{Field: c.Destination.String, Name: "Destination"},
		&validators.TimeIsPresent{Field: c.Lfd.Time, Name: "Lfd"},
		&validators.TimeIsPresent{Field: c.ReservationTime.Time, Name: "ReservationTime"},
		&validators.StringIsPresent{Field: c.Size.String, Name: "Size"},
		&validators.StringIsPresent{Field: c.Type.String, Name: "Type"},
		&validators.StringIsPresent{Field: c.Status.String, Name: "Status"},
		&validators.StringIsPresent{Field: c.GpsURL.String, Name: "GpsURL"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Container) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Container) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
