package util

import "net/http"

func GetQuery(r *http.Request, params string) string {
	return r.URL.Query().Get(params)
}
