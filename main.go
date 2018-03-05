package main

import (
	"log"

	"github.com/omgitsotis/shelter-task/shelter"
)

func main() {
	log.Fatal(shelter.ServeAPI(
		"https://apigateway.test.lifeworks.com/rescue-shelter-api",
	))
}
