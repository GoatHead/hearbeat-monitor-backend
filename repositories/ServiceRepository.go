package repositories

import (
	"goathead/heartbeat-monitor-backend/core/database"
	"goathead/heartbeat-monitor-backend/core/models"
)

func GetService(service *models.Service) (*[]models.Service, error){
	db, _ := database.GetDbConnector()

	var services []models.Service

	query := ""

	if service == nil {
		query = ` SELECT * FROM service s `
	} else {
		query = ` SELECT * FROM service s
				WHERE id = :id`
	}

	err := db.Select(&services, query)

	db.Close()

	return &services, err
}

func AddService(service *models.Service) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` INSERT INTO service (url, name, status) VALUES (:url, :name, :status) `
	tx.NamedExec(query, &service)
	err := tx.Commit()

	return err
}

func UpdateService(service *models.Service) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` UPDATE service SET url = :url,
							name = :name,
                    		status = :status,
							 update_dt = current_timestamp
				WHERE id = :id`
	tx.NamedExec(query, &service)
	err := tx.Commit()

	return err
}

func DeleteService(service *models.Service) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` DELETE FROM service WHERE id = :id `
	tx.NamedExec(query, &service)
	err := tx.Commit()

	return err
}