package service

import (
    "fmt"
    "log"
    "github.com/Petr0752/test/config"
    "github.com/gofiber/fiber/v2"
)

type Service struct {
    cfg *config.Config
}

func NewService(cfg *config.Config) *Service {
    return &Service{cfg: cfg}
}

func (s *Service) Run() {
    app := fiber.New()

    // Определение маршрута GET
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString(s.cfg.SomeField)
    })

    fmt.Println("Сервер запущен на порту 8080")
    log.Fatal(app.Listen(":8080"))
}
