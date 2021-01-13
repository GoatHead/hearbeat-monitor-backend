package webclient

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)


var client *http.Client

func getClient() *http.Client {
	if client == nil {
		timeout:= 10 * time.Second
		defaultTransport := &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client = &http.Client{
			Transport: defaultTransport,
			Timeout: timeout,
		}
	}
	client.CloseIdleConnections()
	return client
}

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

	client := getClient()

	res, err := client.Do(request)

	if err != nil {
		resultCode = -1
	}

	if res != nil {
		resultCode = res.StatusCode
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}

	client.CloseIdleConnections()
	return resultCode
}

func Post(url string, body string) {
	reader := strings.NewReader(body)
	request, _ := http.NewRequest("POST", url, reader)

	client := getClient()

	resp, _ := client.Do(request)

	io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close()

	client.CloseIdleConnections()
}