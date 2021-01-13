package webclient

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
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

	keepAliveTimeout:= 600 * time.Second
	timeout:= 10 * time.Second
	defaultTransport := &http.Transport{
		Dial: (&net.Dialer{
			KeepAlive: keepAliveTimeout,
		}).Dial,
		MaxIdleConns: 100,
		MaxIdleConnsPerHost: 100,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: defaultTransport,
		Timeout: timeout,
	}
	defer client.CloseIdleConnections()

	res, err := client.Do(request)

	if err != nil {
		resultCode = -1
	}

	if res != nil {
		defer res.Body.Close()
		resultCode = res.StatusCode
		_, _ = io.Copy(ioutil.Discard, res.Body)
	}

	return resultCode
}

func Post(url string, body string) {
	reader := strings.NewReader(body)
	request, _ := http.NewRequest("POST", url, reader)
	keepAliveTimeout:= 600 * time.Second
	timeout:= 10 * time.Second
	defaultTransport := &http.Transport{
		Dial: (&net.Dialer{
			KeepAlive: keepAliveTimeout,
		}).Dial,
		MaxIdleConns: 100,
		MaxIdleConnsPerHost: 100,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: defaultTransport,
		Timeout: timeout,
	}
	defer client.CloseIdleConnections()

	resp, _ := client.Do(request)
	defer resp.Body.Close()
	_, _ = io.Copy(ioutil.Discard, resp.Body)

}