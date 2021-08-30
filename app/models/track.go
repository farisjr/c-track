package models

import (
	"database/sql"
	"time"
)

type Track struct {
	TrackID   int `gorm:"primaryKey; unique; not null" json:"track_id" form:"track_id"`
	PatientID int `json:"patient_id" form:"patient_id"`
	CheckerID int `json:"checker_id" form:"checker_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
