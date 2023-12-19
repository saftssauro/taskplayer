package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"goji.io"
	"goji.io/pat"
)

type Fake struct {
	Name string
}

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	ct := r.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single JSON object"
		return &malformedRequest{status: http.StatusBadRequest, msg: msg}
	}

	return nil
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user Fake

	err := decodeJSONBody(w, r, &user)

	if err != nil {
		var badRequest *malformedRequest
		if errors.As(err, &badRequest) {
			http.Error(w, badRequest.msg, badRequest.status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	fmt.Printf("Task: %s", user.Name)
}

func getReportsByUserId(w http.ResponseWriter, r *http.Request) {
	userId := pat.Param(r, "userId")
	print("Pegou o user de id %s", userId)
}

func getTaskByReportId(w http.ResponseWriter, r *http.Request) {
	reportId := pat.Param(r, "reportId")
	print("Get taks do report %s", reportId)
}

func createReport(w http.ResponseWriter, r *http.Request) {
	var report Fake

	err := decodeJSONBody(w, r, &report)

	if err != nil {
		var badRequest *malformedRequest
		if errors.As(err, &badRequest) {
			http.Error(w, badRequest.msg, badRequest.status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	fmt.Printf("Report %s", report.Name)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Fake

	err := decodeJSONBody(w, r, &task)

	if err != nil {
		var badRequest *malformedRequest
		if errors.As(err, &badRequest) {
			http.Error(w, badRequest.msg, badRequest.status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	fmt.Printf("Task: %s", task.Name)
}

func completeTask(w http.ResponseWriter, r *http.Request) {
}

func addRoutes(mux *goji.Mux) {
	mux.HandleFunc(pat.Post("/user/"), createUser)
	mux.HandleFunc(pat.Get("/reports/:userId/"), getReportsByUserId)
	mux.HandleFunc(pat.Post("/reports/"), createReport)
	mux.HandleFunc(pat.Get("/reports/:reportId/tasks/"), getTaskByReportId)
	mux.HandleFunc(pat.Post("/tasks/"), createTask)
	mux.HandleFunc(pat.Patch("/tasks/"), completeTask)
}

func main() {
	root := goji.NewMux()
	addRoutes(root)

	http.ListenAndServe("localhost:8000", root)
}
