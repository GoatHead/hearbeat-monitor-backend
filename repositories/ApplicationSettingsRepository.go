package repositories

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/database"
	"goathead/heartbeat-monitor-backend/core/models"
	"strconv"
)

func GetApplicationSettings(setting *models.ApplicationSettings) (*models.ApplicationSettings, error){
	logger := gin.DefaultWriter
	db, _ := database.GetDbConnector()

	var settings[] models.ApplicationSettings

	query := ""
	param := "param: "
	var err error
	if setting == nil {
		query = ` SELECT * FROM application_settings h `
		err = db.Select(&settings, query)
	} else {
		query = ` SELECT * FROM application_settings h
				WHERE id = ?`
		err = db.Select(&settings, query, setting.Id)
		param += "id=" + strconv.Itoa(setting.Id)
	}

	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	db.Close()

	return &settings[0], err
}

func UpdateApplicationSettings(setting *models.ApplicationSettings) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` UPDATE application_settings SET cycleSec = :cycleSec
				WHERE id = :id`
	tx.NamedExec(query, &setting)

	param := "param: "
	param +=  "id=" + strconv.Itoa(setting.Id)
	param +=  "; cycleSec=" + strconv.Itoa(setting.CycleSec)
	param +=  "\n"
	err := tx.Commit()

	logger := gin.DefaultWriter

	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	return err
}