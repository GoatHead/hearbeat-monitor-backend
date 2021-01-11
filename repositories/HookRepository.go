package repositories

import (
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/database"
	"goathead/heartbeat-monitor-backend/core/models"
	"strconv"
)

func GetHook(hook *models.Hook) (*[]models.Hook, error){
	logger := gin.DefaultWriter
	db, _ := database.GetDbConnector()

	var hooks []models.Hook

	query := ""
	param := "param: "
	var err error
	if hook == nil {
		query = ` SELECT * FROM hook h `
		err = db.Select(&hooks, query)
	} else {
		query = ` SELECT * FROM hook h
				WHERE id = ?`
		err = db.Select(&hooks, query, hook.Id)
		param += "id=" + strconv.Itoa(hook.Id)
	}

	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	db.Close()

	return &hooks, err
}

func AddHook(hook *models.Hook) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` INSERT INTO hook (url, name) VALUES (:url, :name) `
	tx.NamedExec(query, &hook)
	err := tx.Commit()

	param := "param: "
	param +=  "url=" + hook.Url
	param +=  "; name=" + hook.Name
	param +=  "\n"

	logger := gin.DefaultWriter

	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	return err
}

func UpdateHook(hook *models.Hook) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` UPDATE hook SET url = :url,
							name = :name,
							 update_dt = datetime('now', 'localtime')
				WHERE id = :id`
	tx.NamedExec(query, &hook)

	param := "param: "
	param +=  "id=" + strconv.Itoa(hook.Id)
	param +=  "; url=" + hook.Url
	param +=  "; name=" + hook.Name
	param +=  "\n"
	err := tx.Commit()

	logger := gin.DefaultWriter

	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	return err
}

func DeleteHook(hook *models.Hook) error {
	db, _ := database.GetDbConnector()

	tx := db.MustBegin()
	query := ` DELETE FROM hook WHERE id = :id `
	tx.NamedExec(query, &hook)
	err := tx.Commit()

	param := "param: "
	param +=  "id=" + strconv.Itoa(hook.Id)
	param +=  "\n"

	logger := gin.DefaultWriter

	logger.Write([]byte("query:" + query + "\n"))
	logger.Write([]byte(param + "\n"))

	return err
}