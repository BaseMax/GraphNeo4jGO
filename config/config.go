package config

import "time"

type (
	Config struct {
		Postgres Postgres
		Secrets  Secrets
		Server   Server
	}

	Postgres struct {
		URI     string
		Timeout uint16
	}

	Secrets struct {
		JwtSecret string
		ExpTime   time.Duration
	}

	Server struct {
		Addr string
	}
)
