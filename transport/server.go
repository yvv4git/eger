package transport

import (
	"fmt"
	gorillaMux "github.com/gorilla/mux"
	"github.com/yvv4git/eger/internal/config"
	"net/http"
)

// NewWebServer - simple factory for create web server instance
func NewWebServer(webConfig config.WebSrv) *http.Server {
	router := gorillaMux.NewRouter()
	router.HandleFunc("/api/v1/jaeger", jaegerHandler).Methods(http.MethodGet)

	// настройка веб-сервера
	webSrv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", webConfig.Host, webConfig.Port),
		Handler: router,
	}

	return webSrv
}
