package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sekolahkita/go-api/server/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := utils.GetQuery(r, "id")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Hello World from Go! 👋"
	resp["id"] = params
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
	}
}
