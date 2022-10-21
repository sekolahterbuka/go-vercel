package handler

import (
	"net/http"
)

func init() {

}

func Chi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!sssss"))
}
