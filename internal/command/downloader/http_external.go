package downloader

import (
	"net/http"
)

type HttpExternal struct {
}

func (e HttpExternal) GetWithHeader(url string) (http.Response, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Referer", "https://book.pep.com.cn/")
	request.Header.Set("Cookie", "")
	response, err := client.Do(request)
	return *response, err
}
