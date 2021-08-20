package models

type DoctorTestDetails struct {
	Test_ID             int    `gorm:"not null" json:"test_id"`
	Practice_License_ID int    `gorm:"unique;not null" json:"strid"`
	Tests                Tests   `gorm:"foreignkey:ID"`
	Practice_License    Doctor `gorm:"foreignkey:ID"`
}
