package services

import (
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/repositories"
)

func GetServiceHistory(searchCondition *models.SearchCondition) (*[]models.Service, error) {
	service, err := repositories.GetServiceHistory(searchCondition)
	return service, err
}

func GetServiceHistoryCnt(searchCondition *models.SearchCondition) (int, error) {
	cnt, err := repositories.GetServiceHistoryCnt(searchCondition)
	return cnt, err
}