package api

import (
	"fmt"
	"net/http"

	"github.com/gesia-platform/core/api/handler"
)

type API struct {
	server *http.Server
}

func NewAPI(port uint, apiHandler *handler.APIHandler) API {
	api := API{}

	api.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: apiHandler.Mux,
	}

	return api
}

func (api *API) Start() {
	if err := api.server.ListenAndServe(); err != nil {
		panic(err)
	}
}
