package config

import "github.com/kelseyhightower/envconfig"

func ParseEnv(cfg *Config) {
	envconfig.MustProcess("postgres", &cfg.Postgres)
	envconfig.MustProcess("secrets", &cfg.Secrets)
	envconfig.MustProcess("server", &cfg.Server)
	envconfig.MustProcess("neo4j", &cfg.Neo4j)
}
