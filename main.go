package main

import (
    "fmt"
    "log"
    "github.com/Petr0752/test/config"
    "github.com/Petr0752/test/internal/service"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Ошибка загрузки конфигурации: %v", err)
    }

    srv := service.NewService(cfg)
    srv.Run()
}
