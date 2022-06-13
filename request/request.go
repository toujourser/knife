package request

import (
	"bytes"
	"net/http"
)

func Request(method, url string, headers map[string]string, bodyBytes []byte) (resp *http.Response, err error) {
	buffer := bytes.NewBuffer(bodyBytes)

	request, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return
	}
	for key, value := range headers {
		if key == "Host" {
			request.Host = value
		} else {
			request.Header.Set(key, value)
		}
	}
	client := http.Client{}
	resp, err = client.Do(request)
	if err != nil {
		return
	}
	return
}
