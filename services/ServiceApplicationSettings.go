package services

import (
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/repositories"
)

func GetApplicationSettings(setting *models.ApplicationSettings) (*models.ApplicationSettings, error) {
	services, err := repositories.GetApplicationSettings(setting)
	return services, err
}

func UpdateApplicationSettings(setting *models.ApplicationSettings) error {
	err := repositories.UpdateApplicationSettings(setting)
	return err
}