package database

import (
	"app/config"
	"app/models"
)

func CreateOrganization(organization models.Organization) (interface{}, error) {
	if err := config.DB.Save(&organization).Error; err != nil {
		return nil, err
	}
	return organization, nil
}
