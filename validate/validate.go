package validate

import (
	"github.com/asaskevich/govalidator"
)

func ValidateErrorWrapper(err error) map[string]map[string]string {
	errors := map[string]map[string]string{}
	errors["errors"] = govalidator.ErrorsByField(err)
	return errors
}
