package main

import (
	"fmt"
	"legitt-backend/helper"
	"net/http"
)

func GeneratePublicContractWithGptApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var data chatToContractRequest
	helper.DecodeReqJsonToMap(r.Body, &data)

	res, err := handleGeneratePublicContractWithGptApi(data)

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	fmt.Fprint(w, res)
}
