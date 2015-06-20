package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleGreeting(w http.ResponseWriter, r *http.Request) {
	response := config.Respond

	log.Println("-->> Responding: ", response)

	w.Write([]byte(response))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", HandleGreeting)

	http.Handle("/", router)

	err := http.ListenAndServe(config.BindAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
