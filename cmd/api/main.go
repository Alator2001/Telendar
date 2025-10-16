package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	cfgpkg "github.com/Alator2001/Telendar/internal/config"
)

func main() {
	_ = godotenv.Load()

	// 1) Грузим конфиг из окружения
	cfg := cfgpkg.Load()

	// 2) Ставим TZ процесса (влияет на локальное время и логи)
	_ = os.Setenv("TZ", cfg.TZ)
	// 1. Создаём новое приложение Fiber (это как http.Server)
	app := fiber.New()

	// 2. Простейший маршрут — /health
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"ok":  true,
			"env": cfg.Env,
			"tz":  cfg.TZ,
		})
	})

	log.Printf("listening on :%s (env=%s, tz=%s)", cfg.Port, cfg.Env, cfg.TZ)
	if err := app.Listen(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
