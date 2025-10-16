package config

import (
	"log"
	"os"
)

// Config — всё, что нужно приложению на старте.
// Сейчас мало полей, позже добавим DB, почту и пр.
type Config struct {
	Env  string // dev|prod|test
	Port string // порт HTTP-сервера
	TZ   string // таймзона процесса (для логов/дат)
}

// Load читает переменные окружения и заполняет конфиг.
func Load() Config {
	return Config{
		Env:  get("APP_ENV", "dev"),
		Port: get("APP_PORT", "8080"),
		TZ:   get("TZ", "UTC"),
	}
}

func get(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	// дефолт логируем один раз, чтобы понимать, что значение не задано
	if def == "" {
		log.Printf("env %s is empty; using \"\"", key)
	}
	return def
}
