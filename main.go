package main

import (
	"log"
)

func main() {
	err := cli()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println()
}
