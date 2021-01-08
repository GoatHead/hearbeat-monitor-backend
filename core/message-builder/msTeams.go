package message_builder

import (
	"fmt"
	"goathead/heartbeat-monitor-backend/core/models"
	"time"
)

func Build(serviceList []models.Service) string {
	t := time.Now()
	layout := "2006-01-02 15:04:05"
	timeString := t.Format(layout)

	title := fmt.Sprintf("[%s] 하트비트 검사 결과", timeString)
	format := `
		{
            "@type": "MessageCard",
			"@context": "https://schema.org/extensions",
			"summary": "%s",
			"title": "%s",
			"sections": [
				%s
     			]
		}
	`
	sections := ""

	lastIdx := len(serviceList) - 1
	for idx, service := range serviceList {
		sections += buildSection(service)
		if idx != lastIdx {
			sections += ", "
		}
	}

	return fmt.Sprintf(format, title, title, sections)
}

func buildSection(service models.Service) string {
	// url, name, 상태 컬러, 정상/이상, 상태컬러, 스테이터스 코드 순
	format := `{ "text": "<a href=\"%s\"><strong>%s</strong></a> <strong style=\"color: %s\">%s</strong>(<span style=\"color: %s\">%d</span>)" }`
	name := service.Name
	url := service.Url
	statusCode := service.Status

	color := getStatusColor(statusCode)
	statusString := getStatusString(statusCode)

	return fmt.Sprintf(format, url, name, color, statusString, color, statusCode)
}

func getStatusColor(status int) string {
	var result string
	switch switchStatus := status; {
	case switchStatus >= 200 && switchStatus < 300:
		result = "#73BA36"
	case switchStatus >= 300 && switchStatus < 400:
		result = "#FFE731"
	case switchStatus >= 400 && switchStatus < 500:
		result = "#E4636A"
	case switchStatus >= 500 && switchStatus < 600:
		result = "#CF0006"
	case switchStatus == -1:
		result = "red"
	default:
		result = "#D296D2"
	}
	return result
}

func getStatusString(status int) string {
	var result string
	switch switchStatus := status; {
	case switchStatus == 200:
		result = "정상"
	case switchStatus >= 200 && switchStatus < 300:
		result = "일부 성공"
	case switchStatus >= 300 && switchStatus < 400:
		result = "리다이렉션"
	case switchStatus == 400:
		result = "잘못된 요청"
	case switchStatus == 401:
		result = "권한 없음"
	case switchStatus == 404:
		result = "존재하지 않는 리소스"
	case switchStatus >= 400 && switchStatus < 500:
		result = "요청 오류"
	case switchStatus >= 500 && switchStatus < 600:
		result = "서버 오류"
	case switchStatus == -1:
		result = "존재하지 않는 도메인"
	default:
		result = "기타 이상"
	}
	return result
}
