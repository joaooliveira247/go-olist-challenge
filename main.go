package main

import (
	"log"
	"os"

	"github.com/joaooliveira247/go-olist-challenge/src/cmd"
	"github.com/joaooliveira247/go-olist-challenge/src/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	cli := cmd.Gen()

	if err := cli.Run(os.Args); err !=nil {
		log.Fatal(err)
	}
}
