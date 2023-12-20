package validators

import (
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/saftssauro/taskplayer/infra/http/dtos"
	"github.com/saftssauro/taskplayer/infra/http/helpers"
)

type CreateReportBodyValidator struct{}

func (createReportBodyValidator CreateReportBodyValidator) Validate(body io.ReadCloser) (dtos.CreateReportBody, error) {
	dto := dtos.CreateReportBody{}

	helpers.DecodeJSONBody(body, &dto)

	validate := validator.New()
	err := validate.Struct(dto)
	if err != nil {
		fmt.Println(err.Error())

		return dto, err
	}

	return dto, nil
}
