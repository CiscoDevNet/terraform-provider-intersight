package intersight

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// IntBetween returns a SchemaValidateFunc which tests if the provided value
// is of type int and is between min and max (exclusive)
func IntBetweenExclusive(min, max int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v <= min || v >= max {
			errors = append(errors, fmt.Errorf("expected %s to be in the range (%d - %d) (exclusive), got %d", k, min, max, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

// IntAtLeast returns a SchemaValidateFunc which tests if the provided value
// is of type int and is exclusive min
func IntAtLeastExclusive(min int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v <= min {
			errors = append(errors, fmt.Errorf("expected %s to be at least (%d) (exlusive minimum), got %d", k, min+1, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

// IntAtMost returns a SchemaValidateFunc which tests if the provided value
// is of type int and is exclusive of max
func IntAtMostExclusive(max int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v >= max {
			errors = append(errors, fmt.Errorf("expected %s to be at most (%d) (exclusive maximum), got %d", k, max-1, v))
			return warnings, errors
		}

		return warnings, errors
	}
}
