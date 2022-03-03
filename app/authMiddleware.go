package app

import (
	"github.com/Dontunee/banking/domain"
	"github.com/Dontunee/banking/errs"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	repository domain.IAuthRepository
}

func (authMiddleware AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			currentRoute := mux.CurrentRoute(request)
			currentRouteVars := mux.Vars(request)
			authHeader := request.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				isAuthorized := authMiddleware.repository.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

				if isAuthorized {
					//pass to next middleware
					next.ServeHTTP(writer, request)
				} else {
					appError := errs.AppError{Code: http.StatusForbidden, Message: "Unauthorized"}
					writeResponse(writer, appError.Code, appError.AsMessage())
				}
			} else {
				writeResponse(writer, http.StatusUnauthorized, "missing token")
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		//return token itself
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
