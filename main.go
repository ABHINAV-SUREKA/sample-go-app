package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	for {
		link := os.Getenv("LINK")
		resp, err := http.Get("http://" + link + "/api/v1/status/tsdb")
		if err != nil {
			log.Printf("Error: %v", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error: %v", err)
		}

		log.Printf(string(body))

		time.Sleep(20 * time.Second)
	}
}
