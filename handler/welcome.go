package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Welcome(resp http.ResponseWriter, req *http.Request) {

	resp.WriteHeader(http.StatusOK)

	message := fmt.Sprintf("vivek is so cool")

	err := json.NewEncoder(resp).Encode(message)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
	}
}
