package main

import (
	"embed"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"slices"
	"strings"
)

type link struct {
	Name        string
	Description string
	URL         string
	Color       string
	Icon        string
	IconStyle   template.CSS
	Sensitive   bool
}

type page struct {
	Title       string
	BeforeLinks template.HTML
	AfterLinks  template.HTML
	Links       []link
}

//go:embed static/*
var assets embed.FS

//go:embed links/*
var linksFolder embed.FS

//go:embed index.html
var indexFile string

func getLinks() map[string]page {
	linkFiles, err := linksFolder.ReadDir("links")
	if err != nil {
		panic(err)
	}

	links := make(map[string]page)
	for _, e := range linkFiles {
		data, err := os.ReadFile("links/" + e.Name())
		if err != nil {
			panic(err)
		}

		var l page
		err = yaml.Unmarshal(data, &l)
		if err != nil {
			panic(err)
		}
		links[strings.TrimSuffix(e.Name(), ".yml")] = l
	}

	return links
}

func main() {
	darkmodeURL, ok := os.LookupEnv("DARKMODE_URL")
	if !ok {
		darkmodeURL = "https://darkmode.datasektionen.se/"
	}

	links := getLinks()

	indexTemplate, err := template.New("index.html").Parse(indexFile)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get(darkmodeURL)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			slog.Error("Failed fetching darkmode status", "error", err)
		}
		var darkmode bool
		if err := json.NewDecoder(res.Body).Decode(&darkmode); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			slog.Error("Failed parsing darkmode status", "error", err)
		}

		whichLinks := strings.Split(r.Host, ".")[0]
		linksToShow, ok := links[whichLinks]
		if !ok {
			linksToShow = links["systems"]
		}
		if darkmode {
			linksToShow.Links = slices.DeleteFunc(linksToShow.Links, func(l link) bool {
				return l.Sensitive
			})
		}

		if err := indexTemplate.Execute(w, linksToShow); err != nil {
			http.Error(w, "Failed rendering template", http.StatusInternalServerError)
			slog.Error("Failed rendering template", "error", err)
		}
	})

	fs, err := fs.Sub(assets, "static")
	if err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer(http.FS(fs)))

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}
	slog.Info("Started", "port", port)
	http.ListenAndServe(":"+port, nil)
}
