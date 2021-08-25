package models

import (
	"database/sql"
	"time"
)

type TestCategories struct {
	TestCategoriesID int    `gorm:"primaryKey; unique; not null" json:"testcategories_id"`
	Name             string `gorm:"type:char(16); not null" json:"name"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime `gorm:"index"`
}

type TestCategoriesResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    []TestCategories `json:"data"`
}

type TestCategoriesResponseSingle struct {
	Status  bool           `json:"status"`
	Message string         `json:"message"`
	Data    TestCategories `json:"data"`
}
