package main

import (
	"encoding/json"
	"net/http"
)

func (app *Config) Broker(writer http.ResponseWriter, req *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "HEY",
	}

	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		panic(err.Error())
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(out)
}
