package main

import (
    "log"
    "github.com/Petr0752/test/config"
    "github.com/Petr0752/test/internal/service"
	"github.com/joho/godotenv"
)

func main() {
    // Загружаем переменные из .env файла
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Ошибка загрузки файла .env: %v", err)
    }

    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Ошибка загрузки конфигурации: %v", err)
    }

    srv := service.NewService(cfg)
    srv.Run()
}