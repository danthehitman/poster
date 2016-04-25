package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"controller"
	"api"
	"model"
	"config"
	"github.com/kardianos/osext"
	//"path/filepath"
	"fmt"
	//"github.com/gorilla/mux"
)

func main() {
	config.Init()
	//model.InitiDb("user=poster password=admin dbname=poster sslmode=disable search_path=api")
	model.InitiDb(config.Database.ConnectionString, true)
	model.DevStrap{}.InitDevDbRecords()
	log.Println("Database connection complete.")
	templateCache, _ := buildTemplateCache()
	controller.Setup(templateCache)
	api.Init()

	http.ListenAndServe(":8000", nil)
	//go http.ListenAndServe(":8000", nil)

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
	wd, _ := os.Getwd()
	//err := filepath.Walk(wd, visit)
	f, err := os.Open(config.ResourcePrefix + "templates")
	if err != nil{
		log.Println(wd)
		log.Println(osext.ExecutableFolder())
		panic(err)
	}

	fileInfos, _ := f.Readdir(-1)
	fileNames := make([]string, len(fileInfos))
	for idx, fi := range fileInfos {
		if fi.ModTime().After(lastModTime) {
			lastModTime = fi.ModTime()
			needUpdate = true
		}
		fileNames[idx] = config.ResourcePrefix + "templates/" + fi.Name()
	}

	var tc *template.Template
	if needUpdate {
		log.Print("Template change detected, updating...")
		tc = template.Must(template.ParseFiles(fileNames...))
		log.Println("template update complete")
	}
	return tc, needUpdate
}

//func visit(path string, f os.FileInfo, err error) error {
//	fmt.Printf("Visited: %s\n", path)
//	return nil
//}