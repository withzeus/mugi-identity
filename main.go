package main

import (
	"embed"
	"log"

	"github.com/withzeus/mugi-identity/config"
	"github.com/withzeus/mugi-identity/server"
)

//go:embed static/*
var StaticAssets embed.FS

//go:embed templates/*
var TemplateAssets embed.FS

func main() {

	cfg := server.ServerConfig{
		Port:           config.AppEnv.Port,
		TemplateAssets: TemplateAssets,
		StaticAssets:   StaticAssets,
	}
	server := server.New(cfg)

	// html := html.New()
	// for _, temp := range html.GetFS().Templates() {
	// 	log.Printf("%s", temp.Name())
	// }

	log.Printf("Server started on %s \n", config.AppEnv.Port)
	log.Fatal(server.ListenAndServe())
}

// func main() {
// 	staticFiles, err := fs.Sub(StaticAssets, "static")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	staticFileServer := http.FileServerFS(staticFiles)

// 	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

// 	http.HandleFunc("/sign-in", signInHandler)

// 	log.Println("Server started on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
