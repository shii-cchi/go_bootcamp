package app

import (
	"day03/internal/db"
	"day03/internal/es_utils"
	"day03/internal/handlers"
	"log"
	"net/http"
)

func RunCreateIndexAndUploadData() {
	es_utils.CreateIndexAndUploadData()
}

func RunServer() {
	client := es_utils.MakeNewEsClient()
	store := db.NewEsStore(client)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { handlers.GetAllPlacesHandlerHTML(w, r, store) })
	http.HandleFunc("/api/places", func(w http.ResponseWriter, r *http.Request) { handlers.GetAllPlacesHandlerJSON(w, r, store) })

	log.Fatal(http.ListenAndServe(":8888", nil))
}
