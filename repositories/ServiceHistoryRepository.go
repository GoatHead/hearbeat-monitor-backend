package repositories

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/database"
	"goathead/heartbeat-monitor-backend/core/models"
	"strconv"
)

func GetServiceHistory(searchCondition *models.SearchCondition) (*[]models.Service, error){
	db, _ := database.GetDbConnector()

	var services []models.Service

	query := ""
	param := "param: "

	query = ` SELECT * FROM service_history s WHERE 1 = 1`
	args := make([]interface{}, 0)

	if searchCondition != nil {

		startDate := searchCondition.StartDate
		endDate := searchCondition.EndDate

		if startDate != "" && endDate != "" {
			query += ` AND create_dt >= ? AND create_dt <= ? || '23:59:59' `
			param += fmt.Sprintf("startDate=%s; endDate=%s", startDate, endDate)
			args = append(args, startDate, endDate)
		}

		status := searchCondition.Status

		if status != 0 {
			query += ` AND status like '%' || ? || '%' `
			param += fmt.Sprintf("status=%d;", status)
			args = append(args, status)
		}

		url := searchCondition.Url

		if  url != "" {
			query += ` AND url like '%' || ? || '%' `
			param += fmt.Sprintf("url=%s;", url)
			args = append(args, url)
		}

		name := searchCondition.Name

		if  name != "" {
			query += ` AND name like '%' || ? || '%' `
			param += fmt.Sprintf("name=%s;", name)
			args = append(args, name)
		}

		param += "\n"
	}

	err := db.Select(&services, query, args...)

	db.Close()

	logger := gin.DefaultWriter
	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	return &services, err
}

func AddServiceHistory(service *models.Service) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` INSERT INTO service_history (url, name, status) VALUES (:url, :name, :status) `
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