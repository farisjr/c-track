package models

type DoctorTestDetails struct {
	TestID             int    `gorm:"not null" json:"test_id"`
	PracticeLicenseID int    `gorm:"unique;not null" json:"practice_license_id"`
	Tests               Tests  `gorm:"foreignkey:ID"`
	PracticeLicense    Doctor `gorm:"foreignkey:ID"`
}
