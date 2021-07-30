package server

import (
	// "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/data", getData).Methods("GET")
	router.HandleFunc("/data", postData).Methods("POST")
	return router
}

func getData(w http.ResponseWriter, r *http.Request) {
    data, err := os.ReadFile(os.Getenv("DATA_PATH"))
    if err != nil {
        log.Printf("Error reading data file: %v", err)
		http.Error(w, "can't read data file", http.StatusBadRequest)
		return
    }

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func postData(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read request body", http.StatusBadRequest)
		return
	}

	err = os.WriteFile(os.Getenv("DATA_PATH"), body, 0644)
	if err != nil {
		log.Printf("Error writing to file")
		http.Error(w, "can't write to file", http.StatusBadRequest)
	}
}