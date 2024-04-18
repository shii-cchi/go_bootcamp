package app

import (
	"day03/internal/db"
	"day03/internal/es_utils"
	"day03/internal/handlers"
	"log"
	"net/http"
)

func RunCreateIndexAndUploadData() {
	es := es_utils.MakeNewEsClient()

	mappingsJSON := db.GetMappingSchema()

	es_utils.CreateIndex(es, mappingsJSON)

	es_utils.UploadData(es)
}

func RunServer() {
	http.HandleFunc("/", handlers.GetAllPlacesHandler)

	log.Fatal(http.ListenAndServe(":8888", nil))
}
