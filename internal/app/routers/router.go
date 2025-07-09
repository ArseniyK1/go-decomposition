package routers

import (
	"decomposition/internal/app/handlers/echo"
	"decomposition/internal/app/handlers/ping"
	"net/http"
)

func Router(router *http.ServeMux) *http.ServeMux {
	router.HandleFunc("GET /ping", ping.PingHandler)
	router.HandleFunc("POST /echo", echo.EchoHandler)
	return router
}
