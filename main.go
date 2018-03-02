package main

import (
	"log"
	"net/http"

	"github.com/omgitsotis/shelter-task/shelter"
)

func main() {
	log.Fatal(shelter.ServeAPI(&http.Client{}))
}
