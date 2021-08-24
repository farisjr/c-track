package models

type CheckerTestDetails struct {
	Employee_ID int     `gorm:"not null" json:"employeeid"`
	Test_id     int     `gorm:"not null" json:"test_id"`
	Employee    Checker `gorm:"foreignkey:ID"`
	Test        Tests   `gorm:"foreignkey:ID"`
}
