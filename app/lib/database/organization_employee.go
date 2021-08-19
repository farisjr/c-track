package database

import (
	"app/config"
	"app/models"
)

func CreateOrganizationEmployee(organization_employee models.Organization_Employee) (interface{}, error) {
	if err := config.DB.Save(&organization_employee).Error; err != nil {
		return nil, err
	}
	return organization_employee, nil
}
