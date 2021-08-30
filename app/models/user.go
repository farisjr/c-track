package models

import (
	"database/sql"
	"time"
)

type User struct {
	UserID    int    `gorm:"primaryKey; unique; not null" json:"user_id" form:"user_id"`
	Password  string `gorm:"type:varchar(100); not null" json:"password" form:"password"`
	Role      Role   `gorm:"not null" json:"role" form:"role"`
	Token     string `json:"token" form:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

type Role string

const (
	patient Role = "Patient"
	doctor  Role = "Doctor"
	checker Role = "Checker"
)
