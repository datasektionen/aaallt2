package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
)

type system struct {
	Name        string
	Description string
	URL         string
	Color       string
	Icon        string
	IconStyle   template.CSS
	Sensitive   bool
}

//go:embed static/*
var assets embed.FS

//go:embed index.html
var indexFile string

func main() {
	darkmodeURL, ok := os.LookupEnv("DARKMODE_URL")
	if !ok {
		darkmodeURL = "https://darkmode.datasektionen.se/"
	}

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
		systemsToShow := systems
		if darkmode {
			systemsToShow = make([]system, 0, len(systems))
			for _, system := range systems {
				if !system.Sensitive {
					systemsToShow = append(systemsToShow, system)
				}
			}
		}
		if err := indexTemplate.Execute(w, systemsToShow); err != nil {
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
