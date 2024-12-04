package server

import "time"

type Config struct {
	Address         string        `env:"HTTP_ADDRESS, default=localhost:8081"`
	ReadTimeout     time.Duration `env:"HTTP_READ_TIMEOUT, default=5s"`
	IdleTimeout     time.Duration `env:"HTTP_IDLE_TIMEOUT, default=120s"`
	WriteTimeout    time.Duration `env:"HTTP_WRITE_TIMEOUT, default=60s"`
	ShutdownTimeout time.Duration `env:"HTTP_SHUTDOWN_TIMEOUT, default=60s"`
}
