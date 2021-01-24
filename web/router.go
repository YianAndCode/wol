package web

import "net/http"

func init() {
	http.HandleFunc("/wake", handler)
}
