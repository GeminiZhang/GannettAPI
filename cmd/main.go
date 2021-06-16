package main

import (
	"cmd/GannettAPI/internal/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.InitializeDatabase()

	r := mux.NewRouter()
	// Handles get all produce and add one or more produces
	r.HandleFunc("/produce", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			jsonProduceList, err := json.Marshal(database.FetchProduce())
			if err != nil {
				fmt.Printf("Failed to marshal produce list, error: %s", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, string(jsonProduceList))
			w.WriteHeader(http.StatusOK)
		}
		if r.Method == http.MethodPost {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Failed to read request body, error: %s", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			produceList := []database.Produce{}
			if err = json.Unmarshal(body, &produceList); err != nil {
				fmt.Printf("F(ailed to unmarshal request body, error: %s", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			for _, produce := range produceList {
				if err = database.AddProduce(&produce); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}
			w.WriteHeader(http.StatusCreated)
		}
	})

	// Handles delete one produce
	r.HandleFunc("/produce/{ProduceCode}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		produceCode := vars["ProduceCode"]

		if r.Method == http.MethodDelete {
			if err := database.DeleteProduce(produceCode); err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusOK)
		}
	})

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
