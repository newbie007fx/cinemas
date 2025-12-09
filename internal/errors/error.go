package errors

type BaseError struct {
	Code    ErrorType `json:"code"`
	Message string    `json:"message"`
}

func (e BaseError) Error() string {
	return e.Message
}

func (e *BaseError) ToError() error {
	return e
}

func (e BaseError) GetStatusCode() int {
	return int(e.Code) / 100
}

type ErrorType uint

func (et ErrorType) New(message string) *BaseError {
	return &BaseError{
		Code:    et,
		Message: message,
	}
}

const (
	ErrorInvalidRequestBody ErrorType = 40001
	ErrorInvalidPathValue   ErrorType = 40002
	ErrorUnauthorize        ErrorType = 40100
	ErrorInvalidPassword    ErrorType = 40101
	ErrorInvalidToken       ErrorType = 40102
	ErrorExpiredToken       ErrorType = 40103
	ErrorActionFobidden     ErrorType = 40300
	ErrorPathNotFound       ErrorType = 40400
	ErrorQueryNoRow         ErrorType = 40401
	ErrorValidation         ErrorType = 42201
	ErrorInternalServer     ErrorType = 50000
	ErrorQueryDatabase      ErrorType = 50001
)
