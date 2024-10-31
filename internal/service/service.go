package service

import (
    "fmt"
    "your_project/config"
)

type Service struct {
    cfg *config.Config
}

func NewService(cfg *config.Config) *Service {
    return &Service{cfg: cfg}
}

func (s *Service) Run() {
    fmt.Println("Значение переменной:", s.cfg.SomeField)
}
