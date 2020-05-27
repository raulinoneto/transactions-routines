package httpadapters

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func parseBody(r *http.Request, output interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, output)
}

func GetClientIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
