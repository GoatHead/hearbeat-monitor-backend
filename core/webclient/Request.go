package webclient

import (
	"net/http"
	"strings"
)

func Request(url string) int {
	reader := strings.NewReader(`
	{
		"Accept": "text/html,application/xhtml+xml,application/xml",
		"Accept-Encoding": "gzip, deflate",
		"Accept-Charset": "ISO-8859-1",
		"User-Agent": "Mozilla/5.0 (Windows NT 6.2; WOW64; rv:19.0) Gecko/20100101 Firefox/19.0"
	}
`)
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

func Post(url string, body string) {
	reader := strings.NewReader(body)
	request, _ := http.NewRequest("POST", url, reader)
	client := &http.Client{}
	client.Do(request)
}