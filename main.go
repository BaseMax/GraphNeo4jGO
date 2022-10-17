package main

import (
	"GraphNeo4jGO/cmd"
	"GraphNeo4jGO/config"
	"github.com/joho/godotenv"
	"log"
)

var cfg = config.Config{}

func init() {
	if err := godotenv.Load("env.env"); err != nil {
		panic(err)
	}
	config.ParseEnv(&cfg)
}

func main() {
	// fmt.Printf("%#v\n\n", cfg)
	if err := cmd.Run(&cfg); err != nil {
		log.Fatal(err)
	}
}
