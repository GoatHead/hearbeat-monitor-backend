package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goathead/heartbeat-monitor-backend/core/models"
	"goathead/heartbeat-monitor-backend/core/webclient"
	"goathead/heartbeat-monitor-backend/services"
	"net/http"
	"time"
)

type IdList struct {
	Id []int `json:"id"`
}

type testResult struct {
	Url string `json:"url"`
	Status int `json:"status"`
	Datetime string `json:"datetime"`
	Name string `json:"name"`
}

func TestList(c *gin.Context) {
	var idList *IdList
	_ = c.BindJSON(&idList)
	var err error
	var serviceList *[]models.Service
	var resultList []testResult

	if idList != nil {
		idListParams := idList.Id
		serviceList, err = services.GetServiceByIdList(idListParams)

		if serviceList != nil {
			resultList = testServiceList(serviceList)
		}


	} else {
		err = errors.New("ID LIST가 빈 값입니다.")
	}
	if err != nil {
		errorLogger := gin.DefaultErrorWriter
		errorLogger.Write([]byte(err.Error() + "\n"))
		c.JSON(http.StatusOK, gin.H{
			"status": "FAIL",
			"data": nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"data": resultList,
		"message": "하트비트 테스트에 성공하였습니다.",
	})
}

func TestAll(c *gin.Context) {
	var err error
	var serviceList *[]models.Service
	serviceList, err = services.GetService(nil)

	var resultList []testResult
	if serviceList != nil {
		resultList = testServiceList(serviceList)
	}

	if err != nil {
		errorLogger := gin.DefaultErrorWriter
		errorLogger.Write([]byte(err.Error() + "\n"))
		c.JSON(http.StatusOK, gin.H{
			"status": "FAIL",
			"data": nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"data": resultList,
		"message": "하트비트 테스트에 성공하였습니다.",
	})
}

func testServiceList(serviceList *[]models.Service) []testResult {
	var resultList []testResult
	if serviceList != nil {
		for _, service := range *serviceList {
			url := service.Url
			status := webclient.Request(url)
			name := service.Name
			t := time.Now()
			layout := "2006-01-02 15:04:05"
			timeString := t.Format(layout)
			elem := testResult{Name: name, Url: url, Status: status, Datetime: timeString}
			resultList = append(resultList, elem)
		}
	}
	return resultList
}