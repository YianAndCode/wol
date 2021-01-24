package web

import (
	"log"
	"net/http"
)

func Start() {
	log.Fatalln(http.ListenAndServe(":9394", nil))
}
