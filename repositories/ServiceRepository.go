package repositories

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/database"
	"goathead/heartbeat-monitor-backend/core/models"
	"strconv"
)

func GetService(service *models.Service) (*[]models.Service, error){
	db, _ := database.GetDbConnector()

	var services []models.Service

	query := ""
	param := "param: "

	if service == nil {
		query = ` SELECT * FROM service s `
	} else {
		query = ` SELECT * FROM service s
				WHERE id = :id`
		param += "id=" + strconv.Itoa(service.Id)
	}

	err := db.Select(&services, query)

	db.Close()

	logger := gin.DefaultWriter
	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	return &services, err
}

func AddService(service *models.Service) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` INSERT INTO service (url, name, status) VALUES (:url, :name, :status) `
	tx.NamedExec(query, &service)
	err := tx.Commit()

	param := "param: "
	param +=  "url=" + service.Url
	param +=  "; name=" + service.Name
	param +=  "; status=" + strconv.Itoa(service.Status)
	param +=  "\n"

	logger := gin.DefaultWriter

	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	return err
}

func UpdateService(service *models.Service) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` UPDATE service SET url = :url,
							name = :name,
							 update_dt = current_timestamp
				WHERE id = :id`
	tx.NamedExec(query, &service)
	err := tx.Commit()

	param := "param: "
	param +=  "id=" + strconv.Itoa(service.Id)
	param +=  "; url=" + service.Url
	param +=  "; name=" + service.Name
	param +=  "\n"

	logger := gin.DefaultWriter

	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	return err
}

func DeleteService(service *models.Service) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` DELETE FROM service WHERE id = :id `
	tx.NamedExec(query, &service)
	err := tx.Commit()

	param := "param: "
	param +=  "id=" + strconv.Itoa(service.Id)
	param +=  "\n"

	logger := gin.DefaultWriter

	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	return err
}