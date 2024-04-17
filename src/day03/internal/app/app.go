package app

import (
	"day03/internal/config"
	"day03/internal/db"
	"day03/internal/es_utils"
	"day03/internal/models"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

func RunCreateIndexAndUploadData() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading the config: %s", err)
	}

	es_utils.CreateIndexAndUploadData(cfg)
}

const htmlTemplate = `<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Places</title>
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1">
</head>

<body>
<h5>Total: {{.Total}}</h5>
<ul>
    {{range .Places}}
    <li>
        <div>{{.Name}}</div>
        <div>{{.Address}}</div>
        <div>{{.Phone}}</div>
    </li>
    {{end}}
</ul>
<a href="/?page={{.Prev}}"{{if .IsFirstPage}} style="display: none"{{end}}>Previous</a>
<a href="/?page={{.Next}}"{{if .IsLastPage}} style="display: none"{{end}}>Next</a>
<a href="/?page={{.Last}}">Last</a>
</body>
</html>`

const limit = 12

type HTMLData struct {
	Total       int
	Places      []models.Place
	Current     int
	Prev        int
	Next        int
	Last        int
	IsFirstPage bool
	IsLastPage  bool
}

func RunServer() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8888", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	client := es_utils.MakeNewEsClient("https://localhost:9200")
	store := db.NewEsStore(client)

	pageStr := r.URL.Query().Get("page")

	if pageStr == "" {
		http.Error(w, "Missing 'page' parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(pageStr)

	if err != nil || page < 1 {
		http.Error(w, fmt.Sprintf("Invalid 'page' value: '%s'", pageStr), http.StatusBadRequest)
		return
	}

	offset := (page - 1) * limit
	places, total, err := store.GetPlaces(limit, offset)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prevPage := page - 1
	nextPage := page + 1
	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	isFirstPage := page == 1
	isLastPage := page == lastPage

	data := HTMLData{
		Total:       total,
		Places:      places,
		Current:     page,
		Prev:        prevPage,
		Next:        nextPage,
		Last:        lastPage,
		IsFirstPage: isFirstPage,
		IsLastPage:  isLastPage,
	}

	tmpl, err := template.New("index").Parse(htmlTemplate)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
