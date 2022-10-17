package config

import "time"

type (
	Config struct {
		Postgres Postgres
		Secrets  Secrets
		Server   Server
		Neo4j    Neo4j
	}

	Postgres struct {
		URI string
		//Timeout uint16
	}

	Secrets struct {
		JwtSecret string
		ExpTime   time.Duration
	}

	Server struct {
		Addr string
	}

	Neo4j struct {
		URI      string
		Realm    string
		Password string
		Username string
	}
)
