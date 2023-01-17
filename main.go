package main

import (
	"log"
	"net/http"
	"os"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		if counter < 100 {
			counter++
			log.Printf("Request: %d", counter)
		} else {
			log.Printf("Request counter have reached 100, exiting.")
			os.Exit(0)
		}
		defer mutex.Unlock()
	})
	http.ListenAndServe(":8080", nil)
}
