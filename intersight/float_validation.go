package intersight

import (
	"fmt"
	"math"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// FloatDivisibleBy returns a SchemaValidateFunc which tests if the provided value
// is of type float64 and is divisible by a given number
func FloatDivisibleBy(divisor float64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(float64)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be float", k))
			return warnings, errors
		}

		if math.Mod(v, divisor) != 0 {
			errors = append(errors, fmt.Errorf("expected %s to be divisible by %.2f, got: %v", k, divisor, i))
			return warnings, errors
		}

		return warnings, errors
	}
}

// FloatBetween returns a SchemaValidateFunc which tests if the provided value
// is of type float64 and is between min and max (exclusive).
func FloatBetweenExclusive(min, max float64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(float64)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be float", k))
			return
		}

		if v <= min || v >= max {
			es = append(es, fmt.Errorf("expected %s to be in the range (%f - %f) (exclusive) , got %.2f", k, min, max, v))
			return
		}

		return
	}
}

// FloatAtLeast returns a SchemaValidateFunc which tests if the provided value
// is of type float and is at least min (exclusive)
func FloatAtLeastExclusive(min float64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(float64)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be float", k))
			return
		}

		if v <= min {
			es = append(es, fmt.Errorf("expected %s to be at least (%f) (exclusive minimum), got %.2f", k, min+1, v))
			return
		}

		return
	}
}

// FloatAtMost returns a SchemaValidateFunc which tests if the provided value
// is of type float and is at most max (exclusive)
func FloatAtMostExclusive(max float64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(float64)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be float", k))
			return
		}

		if v >= max {
			es = append(es, fmt.Errorf("expected %s to be at most (%f) (exclusive maximum), got %.2f", k, max-1, v))
			return
		}

		return
	}
}

// FloatInSlice returns a SchemaValidateFunc which tests if the provided value
// is of type float and matches the value of an element in the valid slice
func FloatInSlice(valid []float64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(float64)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be float", k))
			return warnings, errors
		}

		for _, validInt := range valid {
			if v == validInt {
				return warnings, errors
			}
		}

		errors = append(errors, fmt.Errorf("expected %s to be one of %v, got %.2f", k, valid, v))
		return warnings, errors
	}
}
