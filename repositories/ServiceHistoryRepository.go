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

	query = ` SELECT * FROM service_history s `
	args := make([]interface{}, 0)

	query += getWhereClause(searchCondition, &args, &param)

	if searchCondition == nil {
		query += ` LIMIT 10 OFFSET 0 `
	} else {
		// 페이지네이션
		pageStart := searchCondition.PageStart
		pageSize := searchCondition.PageSize

		if pageStart != 0 && pageSize != 0 {
			query += ` LIMIT ? OFFSET (? - 1) * ? `
			param += fmt.Sprintf("pageSize=%d; pageStart=%d; pageSize=%d;", pageSize, pageStart, pageSize)
			args = append(args, pageSize, pageStart, pageSize)
		} else {
			query += ` LIMIT 10 OFFSET 0 `
		}

	}

	err := db.Select(&services, query, args...)

	db.Close()

	logger := gin.DefaultWriter
	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	return &services, err
}

func GetServiceHistoryCnt(searchCondition *models.SearchCondition) (int, error){
	db, _ := database.GetDbConnector()

	var count int

	query := ""
	param := "param: "

	query = ` SELECT count(1) FROM service_history s `
	args := make([]interface{}, 0)

	query += getWhereClause(searchCondition, &args, &param)

	row := db.QueryRow(query, args...)
	err := row.Scan(&count)

	db.Close()

	logger := gin.DefaultWriter
	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	return count, err
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

func getWhereClause(searchCondition *models.SearchCondition, args *[]interface{}, param *string) string {
	where := "WHERE 1 = 1"

	if searchCondition != nil {

		startDate := searchCondition.StartDate
		endDate := searchCondition.EndDate

		if startDate != "" && endDate != "" {
			where += ` AND create_dt >= ? AND create_dt <= ? || '23:59:59' `
			*param += fmt.Sprintf("startDate=%s; endDate=%s", startDate, endDate)
			*args = append(*args, startDate, endDate)
		}

		status := searchCondition.Status

		if status != 0 {
			where += ` AND status like '%' || ? || '%' `
			*param += fmt.Sprintf("status=%d;", status)
			*args = append(*args, status)
		}

		url := searchCondition.Url

		if  url != "" {
			where += ` AND url like '%' || ? || '%' `
			*param += fmt.Sprintf("url=%s;", url)
			*args = append(*args, url)
		}

		name := searchCondition.Name

		if  name != "" {
			where += ` AND name like '%' || ? || '%' `
			*param += fmt.Sprintf("name=%s;", name)
			*args = append(*args, name)
		}

		*param += "\n"
	}

	return where
}