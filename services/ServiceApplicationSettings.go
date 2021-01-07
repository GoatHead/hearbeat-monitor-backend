package services

import (
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/repositories"
)

func GetApplicationSettings(service *models.ApplicationSettings) (*models.ApplicationSettings, error) {
	services, err := repositories.GetApplicationSettings(service)
	return services, err
}

func UpdateApplicationSettings(service *models.ApplicationSettings) error {
	err := repositories.UpdateApplicationSettings(service)
	return err
}