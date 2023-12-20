package middlewares

import (
	"net/http"
	"strings"
)

func AuthenticateMiddleware(handler http.Handler) http.Handler {
	mux := func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")

		// TO DO: Handle real JWT...
		userId := strings.Split(token, "Bearer ")[1]
		req.Header.Add("userId", userId)

		handler.ServeHTTP(res, req)
	}

	return http.HandlerFunc(mux)
}
