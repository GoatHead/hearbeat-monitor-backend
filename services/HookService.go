package services

import (
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/repositories"
)

func GetHook(hook *models.Hook) (*[]models.Hook, error) {
	hooks, err := repositories.GetHook(hook)
	return hooks, err
}

func AddHook(hook *models.Hook) error {
	err := repositories.AddHook(hook)
	return err
}

func UpdateHook(hook *models.Hook) error {
	err := repositories.UpdateHook(hook)
	return err
}

func DeleteHook(hook *models.Hook) error {
	err := repositories.DeleteHook(hook)
	return err
}