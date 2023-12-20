package dtos

type CreateReportBody struct {
	Name string `json:"name" validate:"required"`
}
