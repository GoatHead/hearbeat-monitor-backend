package services

import (
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/repositories"
)

func GetService(service *models.Service) (*[]models.Service, error) {
	services, err := repositories.GetService(service)
	return services, err
}

func GetServiceByIdList(idList []int) (*[]models.Service, error) {
	services, err := repositories.GetServiceByIdList(idList)
	return services, err
}

func AddService(service *models.Service) error {
	err := repositories.AddService(service)
	return err
}

func UpdateService(service *models.Service) error {
	err := repositories.UpdateService(service)
	return err
}

func DeleteService(service *models.Service) error {
	err := repositories.DeleteService(service)
	return err
}