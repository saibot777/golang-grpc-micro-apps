package main

import (
	"net/http"
)

func (app *Config) Broker(writer http.ResponseWriter, req *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "HEY",
	}

	err := app.writeJson(writer, http.StatusOK, payload)

	if err != nil {
		panic(err.Error())
	}
}
