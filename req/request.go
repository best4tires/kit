package req

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/best4tires/kit/srv"
)

func GetJSON[T any](clt *http.Client, url string) (T, error) {
	var t T
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return t, fmt.Errorf("new-request %q: %w", url, err)
	}
	r.Header.Add(srv.HeaderAccept, srv.ContentTypeJSONUTF8)
	return doJSON[T](clt, r)
}

func PostJSON[T any](clt *http.Client, url string, data any) (T, error) {
	var t T
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return t, fmt.Errorf("json.encode: %w", err)
	}
	r, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return t, fmt.Errorf("new-request %q: %w", url, err)
	}
	r.Header.Add(srv.HeaderContentType, srv.ContentTypeJSONUTF8)
	r.Header.Add(srv.HeaderAccept, srv.ContentTypeJSONUTF8)
	return doJSON[T](clt, r)
}

func doJSON[T any](clt *http.Client, r *http.Request) (T, error) {
	var t T
	resp, err := clt.Do(r)
	if err != nil {
		return t, fmt.Errorf("client.do: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return t, fmt.Errorf("status-code: %s", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return t, fmt.Errorf("json.decode: %w", err)
	}
	return t, nil
}
