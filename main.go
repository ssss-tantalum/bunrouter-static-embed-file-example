package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/ssss-tantalum/bunrouter-static-embed-file-example/internal/app"
	"github.com/ssss-tantalum/bunrouter-static-embed-file-example/internal/app/handler"
	"github.com/uptrace/bunrouter"
)

//go:embed static
var filesFS embed.FS

func main() {
	app := app.New()

	fileServer := http.FileServer(http.FS(filesFS))

	router := app.Router()
	IndexHandler := handler.NewIndexHandler(app)

	router.GET("/", IndexHandler.Index)
	router.GET("/static/*path", bunrouter.HTTPHandler(fileServer))

	log.Println("listening on http://localhost:3000")
	log.Println(http.ListenAndServe(":3000", router))
}
