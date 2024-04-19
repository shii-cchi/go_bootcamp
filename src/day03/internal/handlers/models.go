package handlers

import "day03/internal/db"

const resultsPerPage = 10
const closestPlaces = 3

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

type HTMLData struct {
	Total       int
	Places      []db.Place
	Current     int
	Prev        int
	Next        int
	Last        int
	IsFirstPage bool
	IsLastPage  bool
}

type JSONData struct {
	Name   string     `json:"name"`
	Total  int        `json:"total"`
	Places []db.Place `json:"places"`
	Prev   *int       `json:"prev_page,omitempty"`
	Next   *int       `json:"next_page,omitempty"`
	Last   int        `json:"last_page"`
}
