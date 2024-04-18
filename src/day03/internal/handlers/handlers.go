package handlers

import (
	"day03/internal/db"
	"day03/internal/es_utils"
	"errors"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func GetAllPlacesHandler(w http.ResponseWriter, r *http.Request) {
	client := es_utils.MakeNewEsClient()
	store := db.NewEsStore(client)

	page, err := getPageParam(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset := (page - 1) * limit

	places, total, err := store.GetPlaces(limit, offset)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := prepareHTMLData(page, total, limit, places)

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

func getPageParam(r *http.Request) (int, error) {
	pageStr := r.URL.Query().Get("page")

	if pageStr == "" {
		return 0, errors.New("Missing 'page' parameter")
	}

	page, err := strconv.Atoi(pageStr)

	if err != nil || page < 1 {
		return 0, fmt.Errorf("Invalid 'page' value: '%s'", pageStr)
	}

	return page, nil
}

func prepareHTMLData(page, total, limit int, places []db.Place) HTMLData {
	prevPage := page - 1
	nextPage := page + 1
	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	isFirstPage := page == 1
	isLastPage := page == lastPage

	return HTMLData{
		Total:       total,
		Places:      places,
		Current:     page,
		Prev:        prevPage,
		Next:        nextPage,
		Last:        lastPage,
		IsFirstPage: isFirstPage,
		IsLastPage:  isLastPage,
	}
}
