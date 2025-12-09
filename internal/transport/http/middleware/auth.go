package middleware

import (
	"net/http"
	"strings"

	"github.com/newbie007fx/cinemas/internal/errors"
	"github.com/newbie007fx/cinemas/internal/transport/http/helpers/authentication"
	"github.com/newbie007fx/cinemas/internal/transport/http/helpers/response"
)

const (
	AuthorizationHeader = "Authorization"
)

func Auth(tokenAuth *authentication.TokenAuth) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
			authorization := req.Header.Get(AuthorizationHeader)
			splitAuth := strings.SplitN(authorization, " ", 2)

			if len(splitAuth) != 2 || splitAuth[0] != "Bearer" {
				err := errors.ErrorInvalidToken.New("Invalid or expired credential")
				response.WriterResponseError(wr, err)

				return
			}

			claims, err := tokenAuth.ParseJWT(splitAuth[1])
			if err != nil || claims.TokenType != authentication.ACCESS_TOKEN {
				err := errors.ErrorInvalidToken.New("Invalid or expired credential")
				response.WriterResponseError(wr, err)

				return
			}

			next.ServeHTTP(wr, req)
		})
	}
}
