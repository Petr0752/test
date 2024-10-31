package main

import (
    "fmt"
    "log"
    "test/config"
    "test/internal/service"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Ошибка загрузки конфигурации: %v", err)
    }

    srv := service.NewService(cfg)
    srv.Run()
}
