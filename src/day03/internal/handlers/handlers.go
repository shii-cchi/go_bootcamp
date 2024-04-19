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

func GetAllPlacesHTMLHandler(w http.ResponseWriter, r *http.Request, store db.Store) {
	page, err := getPageParam(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset := (page - 1) * resultsPerPage

	places, total, err := store.GetPlaces(resultsPerPage, offset)

	if err != nil {
		if err.Error() == "Error: 400 Bad Request: all shards failed" {
			http.Error(w, fmt.Sprintf("Invalid 'page' value: '%d'", page), http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := prepareHTMLData(page, total, resultsPerPage, places)

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

	maxDoc, err := es_utils.GetIndexDocCount("places")

	maxPage := int(math.Ceil(float64(maxDoc) / float64(resultsPerPage)))

	if err != nil || page < 1 || page > maxPage {
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

func GetAllPlacesJSONHandler(w http.ResponseWriter, r *http.Request, store db.Store) {
	page, err := getPageParam(r)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	offset := (page - 1) * resultsPerPage

	places, total, err := store.GetPlaces(resultsPerPage, offset)

	if err != nil {
		if err.Error() == "Error: 400 Bad Request: all shards failed" {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid 'page' value: '%d'", page))
			return
		}

		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	data := prepareJSONData(page, total, resultsPerPage, places)

	respondWithJSON(w, http.StatusOK, data)
}

func prepareJSONData(page, total, limit int, places []db.Place) JSONData {
	var prevPage *int

	if page > 1 {
		prev := page - 1
		prevPage = &prev
	}

	lastPage := int(math.Ceil(float64(total) / float64(limit)))

	var nextPage *int

	if page != lastPage {
		next := page + 1
		nextPage = &next
	}

	return JSONData{
		Name:   "Places",
		Total:  total,
		Places: places,
		Prev:   prevPage,
		Next:   nextPage,
		Last:   lastPage,
	}
}

func GetClosestPlacesHandler(w http.ResponseWriter, r *http.Request, store db.Store) {
	lat, lon, err := getLatLonParam(r)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	places, err := store.GetClosestPlaces(lat, lon, closestPlaces)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, places)
}

func getLatLonParam(r *http.Request) (float64, float64, error) {
	latStr := r.URL.Query().Get("lat")

	if latStr == "" {
		return 0, 0, errors.New("Missing 'lat' parameter")
	}

	lat, err := strconv.ParseFloat(latStr, 64)

	if err != nil {
		return 0, 0, fmt.Errorf("Invalid 'lat' value: '%s'", latStr)
	}

	lonStr := r.URL.Query().Get("lon")

	if lonStr == "" {
		return 0, 0, errors.New("Missing 'lon' parameter")
	}

	lon, err := strconv.ParseFloat(lonStr, 64)

	if err != nil {
		return 0, 0, fmt.Errorf("Invalid 'lon' value: '%s'", lonStr)
	}

	return lat, lon, err
}
