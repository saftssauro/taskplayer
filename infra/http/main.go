package http

import (
	"fmt"
	"net/http"

	"github.com/saftssauro/taskplayer/infra/http/routes"
	"goji.io"
)

var API_ADDRESS = "localhost:8000"

func CreateServer(url string) {
	root := goji.NewMux()
	routes.CreateReportsRoutes(root)

	err := http.ListenAndServe(API_ADDRESS, root)
	if err != nil {
		fmt.Println("Error on listen http server:\n", err.Error())
		return
	}
}
