package repositories

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/database"
	"goathead/heartbeat-monitor-backend/core/models"
)

func GetHook(hook *models.Hook) (*[]models.Hook, error){
	logger := gin.DefaultWriter
	logger.Write([]byte("{{}}\n"))
	db, _ := database.GetDbConnector()

	var hooks []models.Hook

	query := ""

	if hook == nil {
		query = ` SELECT * FROM hook h `
	} else {
		query = ` SELECT * FROM hook h
				WHERE id = :id`
	}

	err := db.Select(&hooks, query)

	db.Close()

	return &hooks, err
}

func AddHook(hook *models.Hook) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` INSERT INTO hook (url, name) VALUES (:url, :name) `
	tx.NamedExec(query, &hook)
	err := tx.Commit()

	return err
}

func UpdateHook(hook *models.Hook) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` UPDATE hook SET url = :url,
							name = :name,
							 update_dt = current_timestamp
				WHERE id = :id`
	tx.NamedExec(query, &hook)
	err := tx.Commit()

	return err
}

func DeleteHook(hook *models.Hook) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` DELETE FROM hook WHERE id = :id `
	tx.NamedExec(query, &hook)
	err := tx.Commit()

	return err
}