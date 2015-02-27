package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type group struct {
	Name      string
	Downloads []download
}

type download struct {
	Name, Filename string
}

var files = []group{
	{
		"Wards",
		[]download{
			{"Shapefile", "chicago_2015_precincts.zip"},
			{"TopoJSON", "chicago_2015_precincts_topo.json"},
			{"GeoJSON", "chicago_2015_precincts.json"},
			{"KML", "chicago_2015_precincts.kml"},
		},
	},
	{
		"Precincts",
		[]download{
			{"Shapefile", "chicago_2015_wards.zip"},
			{"TopoJSON", "chicago_2015_wards_topo.json"},
			{"GeoJSON", "chicago_2015_wards.json"},
			{"KML", "chicago_2015_wards.kml"},
		},
	},
	{
		"Combined Wards & Precincts",
		[]download{
			{"PostgreSQL dump", "chicago_2015_wards_precincts.sql.gz"},
		},
	},
}

const baseURL = "https://raw.githubusercontent.com/paulsmith/chicago_wards_and_precincts/2528342e082073fd8055db127eb7330bd77fc3d9/"

func main() {
	funcMap := template.FuncMap{
		"url": func(filename string) string {
			return baseURL + filename
		},
		"slug": func(s string) string {
			return strings.Replace(strings.ToLower(s), " ", "-", -1)
		},
		"unsafe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}
	t := template.Must(template.New("index.html.tmpl").Funcs(funcMap).ParseFiles("index.html.tmpl"))
	if err := t.Execute(os.Stdout, struct {
		Files              []group
		WardsShapefile     string
		PrecinctsShapefile string
	}{
		files,
		baseURL + files[0].Downloads[0].Filename,
		baseURL + files[1].Downloads[0].Filename,
	}); err != nil {
		log.Fatal(err)
	}
}
