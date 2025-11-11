package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

func FornmatErrors(errors *validate.Errors) string {
	var res string

	for key, value := range errors.Errors {
		res = res + key + ": " + strings.Join(value, ", ") + ", "
	}

	return res

}
