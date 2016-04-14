package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"controller"
	"api"
	"model"
	"config"
)

func main() {
	config.Init()
	//model.InitiDb("user=poster password=admin dbname=poster sslmode=disable search_path=api")
	model.InitiDb(config.Database.ConnectionString)
	templateCache, _ := buildTemplateCache()
	controller.Setup(templateCache)
	api.Setup()

	go http.ListenAndServe(":8000", nil)

	go func() {
		for range time.Tick(300 * time.Millisecond) {
			tc, isUpdated := buildTemplateCache()
			if isUpdated {
				controller.SetTemplateCache(tc)
			}
		}
	}()

	log.Println("Server started, press <ENTER> to exit")
	fmt.Scanln()
}

var lastModTime time.Time = time.Unix(0, 0)

func buildTemplateCache() (*template.Template, bool) {
	needUpdate := false

	f, _ := os.Open("templates")

	fileInfos, _ := f.Readdir(-1)
	fileNames := make([]string, len(fileInfos))
	for idx, fi := range fileInfos {
		if fi.ModTime().After(lastModTime) {
			lastModTime = fi.ModTime()
			needUpdate = true
		}
		fileNames[idx] = "templates/" + fi.Name()
	}

	var tc *template.Template
	if needUpdate {
		log.Print("Template change detected, updating...")
		tc = template.Must(template.ParseFiles(fileNames...))
		log.Println("template update complete")
	}
	return tc, needUpdate
}
