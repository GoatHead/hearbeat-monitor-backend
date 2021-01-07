package webclient

import (
	"net/http"
	"strings"
)

func Request(url string) int {
	reader := strings.NewReader(`{"body": ""}`)
	request, _ := http.NewRequest("GET", url, reader)
	var resultCode int

	client := &http.Client{}
	res, err := client.Do(request)

	if err != nil {
		resultCode = -1
	} else {
		resultCode = res.StatusCode
	}

	return resultCode
}