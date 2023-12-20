package routes

import (
	"github.com/saftssauro/taskplayer/infra/http/controllers"
	"github.com/saftssauro/taskplayer/infra/http/middlewares"
	"github.com/saftssauro/taskplayer/infra/http/validators"

	"goji.io"
	"goji.io/pat"
)

func CreateReportsRoutes(mux *goji.Mux) {
	createReportBodyValidator := validators.CreateReportBodyValidator{}
	reportsController := controllers.ReportsController{}.New(createReportBodyValidator)

	reports := goji.SubMux()
	reports.Use(middlewares.AuthenticateMiddleware)

	mux.Handle(pat.New("/reports/*"), reports)
	reports.HandleFunc(pat.Post("/"), reportsController.Create)
	reports.HandleFunc(pat.Get("/"), reportsController.List)
}
