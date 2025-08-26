package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	srv := server.NewServer(logger)

	logger.Println("Запуск сервера...")

	if err := srv.HTTP.ListenAndServe(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
