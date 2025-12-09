package authentication

import (
	"fmt"
	"net/http"

	"github.com/newbie007fx/cinemas/internal/errors"
	"github.com/newbie007fx/cinemas/internal/transport/http/helpers/authentication"
	"github.com/newbie007fx/cinemas/internal/transport/http/helpers/request"
	"github.com/newbie007fx/cinemas/internal/transport/http/helpers/response"
)

func (ah AuthHandlers) Login(rw http.ResponseWriter, req *http.Request) {
	var requestBody authentication.LoginRequst
	err := request.ReadRequestBody(req, &requestBody)
	if err != nil {
		response.WriterResponseError(rw, err)
		return
	}

	userData, err := ah.userUsecase.VerifyUsernamePassword(req.Context(), requestBody.Username, requestBody.Password)
	if err != nil {
		if err.Code == errors.ErrorQueryNoRow || err.Code == errors.ErrorInvalidPassword {
			err = errors.ErrorUnauthorize.New("Invalid username or password")
		}
		response.WriterResponseError(rw, err)
		return
	}

	accessToken, refreshToken, errD := ah.authToken.GenerateAuthToken(userData)
	if errD != nil {
		err = errors.ErrorInternalServer.New(fmt.Sprint("cannot process request: ", errD.Error()))
		response.WriterResponseError(rw, err)
		return
	}

	resp := response.Response[authentication.LoginResponse]{
		IsSuccess: true,
		Data: authentication.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			User:         *userData,
		},
	}

	resp.Send(rw)
}
