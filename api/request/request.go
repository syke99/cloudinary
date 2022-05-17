package request

import (
	"io"
	"net/http"
)

type Request struct {
	request
}

type request interface {
	RequestMedia(string) []byte
}

func (r Request) RequestMedia(client *http.Client, url string) []byte {
	var res []byte

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Content-Type", "application/octet-stream")

	resp, err := client.Do(request)

	if err == nil && resp.StatusCode == 200 {
		defer resp.Body.Close()
		res, err = io.ReadAll(resp.Body)
	}

	return res
}
