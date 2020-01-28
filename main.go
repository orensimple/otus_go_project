package main

import (
	"log"

	"github.com/orensimple/otus_go_project/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
