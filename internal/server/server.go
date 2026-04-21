package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Http   *http.Server
}

func Router(logName *log.Logger) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainHandle)
	mux.HandleFunc("/upload", handlers.DownloaderHandle)

	srv := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logName,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	servStruc := Server{Logger: logName, Http: &srv}

	return &servStruc

}
