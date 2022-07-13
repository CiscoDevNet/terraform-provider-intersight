package intersight

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// StringLenMaximum returns a SchemaValidateFunc which tests if the provided value
// is of type string and has length lesser than max (inclusive)
func StringLenMaximum(max int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(string)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
			return warnings, errors
		}

		if len(v) > max {
			errors = append(errors, fmt.Errorf("expected length of %s to be atmost %d, got %s of length %d", k, max, v, len(v)))
		}

		return warnings, errors
	}
}

// StringLenMinimum returns a SchemaValidateFunc which tests if the provided value
// is of type string and has length greater than min (inclusive)
func StringLenMinimum(min int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(string)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
			return warnings, errors
		}

		if len(v) < min {
			errors = append(errors, fmt.Errorf("expected length of %s to be atleast %d, got %s of length %d", k, min, v, len(v)))
		}

		return warnings, errors
	}
}
