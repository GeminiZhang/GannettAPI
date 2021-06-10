package main

import (
	"cmd/GannettAPI/internal/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	jsonProduceList, _ := json.Marshal(database.FetchProduce())
	fmt.Fprintf(w, string(jsonProduceList))
}

func main() {
	database.InitializeDatabase()
	database.DeleteProduce("A12T-4GH7-QPL9-3N4M")
	fmt.Println(database.ProduceList)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
