// Copyright 2023 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package compare provides a lightweight comparison function on any type.
// reference: https://github.com/stretchr/testify
package compare

import (
	"fmt"
	json "github.com/json-iterator/go"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/exp/constraints"
)

// operator type
const (
	equal          = "eq"
	lessThan       = "lt"
	greaterThan    = "gt"
	lessOrEqual    = "le"
	greaterOrEqual = "ge"
)

var (
	timeType  = reflect.TypeOf(time.Time{})
	bytesType = reflect.TypeOf([]byte{})
)

// Equal checks if two values are equal or not. (check both type and value)
// Play: https://go.dev/play/p/wmVxR-to4lz
func Equal(left, right any) bool {
	return compareValue(equal, left, right)
}

// EqualValue checks if two values are equal or not. (check value only)
// Play: https://go.dev/play/p/fxnna_LLD9u
func EqualValue(left, right any) bool {
	ls, rs := ToString(left), ToString(right)
	return ls == rs
}

// LessThan checks if value `left` less than value `right`.
// Play: https://go.dev/play/p/cYh7FQQj0ne
func LessThan(left, right any) bool {
	return compareValue(lessThan, left, right)
}

// GreaterThan checks if value `left` greater than value `right`.
// Play: https://go.dev/play/p/9-NYDFZmIMp
func GreaterThan(left, right any) bool {
	return compareValue(greaterThan, left, right)
}

// LessOrEqual checks if value `left` less than or equal to value `right`.
// Play: https://go.dev/play/p/e4T_scwoQzp
func LessOrEqual(left, right any) bool {
	return compareValue(lessOrEqual, left, right)
}

// GreaterOrEqual checks if value `left` greater than or equal to value `right`.
// Play: https://go.dev/play/p/vx8mP0U8DFk
func GreaterOrEqual(left, right any) bool {
	return compareValue(greaterOrEqual, left, right)
}

// InDelta checks if two values are equal or not within a delta.
// Play: https://go.dev/play/p/TuDdcNtMkjo
func InDelta[T constraints.Integer | constraints.Float](left, right T, delta float64) bool {
	return float64(Abs(left-right)) <= delta
}

func ToString(value any) string {
	if value == nil {
		return ""
	}

	switch val := value.(type) {
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case int:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case string:
		return val
	case []byte:
		return string(val)
	default:
		b, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		return string(b)

		// todo: maybe we should't supprt other type conversion
		// v := reflect.ValueOf(value)
		// log.Panicf("Unsupported data type: %s ", v.String())
		// return ""
	}
}

// ToFloat convert value to float64, if input is not a float return 0.0 and error.
// Play: https://go.dev/play/p/4YTmPCibqHJ
func ToFloat(value any) (float64, error) {
	v := reflect.ValueOf(value)

	result := 0.0
	err := fmt.Errorf("ToInt: unvalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = float64(v.Int())
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = float64(v.Uint())
		return result, nil
	case float32, float64:
		result = v.Float()
		return result, nil
	case string:
		result, err = strconv.ParseFloat(v.String(), 64)
		if err != nil {
			result = 0.0
		}
		return result, err
	default:
		return result, err
	}
}

// Abs returns the absolute value of x.
// Play: https://go.dev/play/p/fsyBh1Os-1d
func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return (-x)
	}

	return x
}
