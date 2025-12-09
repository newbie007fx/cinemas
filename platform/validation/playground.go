package validation

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var service *ValidatorService

type ValidatorService struct {
	Validator *validator.Validate
}

func New() *ValidatorService {
	service = &ValidatorService{}

	return service
}

func (vs *ValidatorService) Setup() error {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		return name
	})

	vs.Validator = validate

	return nil
}

func Validate(data interface{}) error {
	err := service.Validator.Struct(data)
	if err != nil {
		return translateError(err)
	}

	return nil
}

func translateError(err error) error {
	errMsg := ""

	if v, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range v {
			errMsg = getValidationMsg(fieldError.Field(), fieldError.Tag())
		}
	}

	return errors.New(errMsg)
}

var msgs map[string]string = map[string]string{
	"required": "{name} can not be empty",
}

func getValidationMsg(name, tag string) string {
	msg, ok := msgs[tag]
	if !ok {
		msg = "{name} is invalid"
	}

	return strings.ReplaceAll(msg, "{name}", name)
}
