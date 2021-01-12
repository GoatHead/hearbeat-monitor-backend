package repositories

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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

func GetServiceByIdList(idList []int) (*[]models.Service, error){
	db, _ := database.GetDbConnector()

	var services []models.Service

	query := ""
	param := "param: "

	query, args, _ := sqlx.In("SELECT * FROM service s WHERE id IN (?);", idList)
	param += "ids=" + fmt.Sprint(idList)

	query = db.Rebind(query)
	err := db.Select(&services, query, args...)

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

	db.Close()

	return err
}

func UpdateService(service *models.Service) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` UPDATE service SET url = :url,
							name = :name,
							 update_dt = datetime('now', 'localtime')
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

	db.Close()

	return err
}

func UpdateServiceStatusCode(service *models.Service) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` UPDATE service SET status = :status,
							 update_dt = datetime('now', 'localtime')
				WHERE id = :id`
	tx.NamedExec(query, &service)
	err := tx.Commit()

	param := "param: "
	param +=  "id=" + strconv.Itoa(service.Id)
	param +=  "; status=" + strconv.Itoa(service.Status)
	param +=  "\n"

	logger := gin.DefaultWriter

	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	db.Close()

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

	db.Close()

	return err
}