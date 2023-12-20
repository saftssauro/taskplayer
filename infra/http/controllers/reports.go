package controllers

import (
	"fmt"
	"net/http"

	"github.com/saftssauro/taskplayer/infra/http/validators"
)

type ReportsController struct {
	createReportBodyValidator validators.CreateReportBodyValidator
}

func (reportController ReportsController) New(createReportBodyValidator validators.CreateReportBodyValidator) *ReportsController {
	return &ReportsController{
		createReportBodyValidator: createReportBodyValidator,
	}
}

func (reportController ReportsController) Create(res http.ResponseWriter, req *http.Request) {
	dto, err := reportController.createReportBodyValidator.Validate(req.Body)

	if err != nil {
		res.WriteHeader(400)
		return
	}

	fmt.Println(dto)
	res.WriteHeader(201)
}

func (reportController ReportsController) List(res http.ResponseWriter, req *http.Request) {
	fmt.Println("list")
}
