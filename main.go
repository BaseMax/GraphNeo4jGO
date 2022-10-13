package main

import (
	"GraphNeo4jGO/cmd"
	"GraphNeo4jGO/config"
)

var cfg = config.Config{}

func init() {
    config.ParseEnv(&cfg)
}

func main() {
    if err := cmd.Run(&cfg); err != nil {
        panic(err)
    }
}
