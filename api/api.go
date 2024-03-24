package api

import (
	"fmt"
	"net/http"

	"github.com/gesia-platform/core/api/handler"
	"github.com/gesia-platform/core/context"
)

type API struct {
	server *http.Server
}

func NewAPI(port uint, context *context.Context) API {
	api := API{}

	handler := handler.NewAPIHandler(context)

	api.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler.Mux,
	}

	return api
}

func (api *API) Start() {
	if err := api.server.ListenAndServe(); err != nil {
		panic(err)
	}
}
