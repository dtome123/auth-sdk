package jwtutils

import (
	"fmt"
	"reflect"
	"strconv"
)

type Value struct {
	Value interface{}
}

// Raw returns raw underlying value.
func (v Value) Raw() interface{} {
	return v.Value
}

// IsNil returns true if the value is nil.
func (v Value) IsNil() bool {
	return v.Value == nil
}

// IsZero returns true if value is nil or empty string.
func (v Value) IsZero() bool {
	return v.Value == nil || v.AsString() == ""
}

// AsString converts value to string.
func (v Value) AsString() string {
	switch val := v.Value.(type) {
	case string:
		return val
	case float64, int, int64, bool:
		return fmt.Sprintf("%v", val)
	default:
		return ""
	}
}

// AsInt converts value to int.
func (v Value) AsInt() int {
	switch val := v.Value.(type) {
	case int:
		return val
	case float64:
		return int(val)
	case string:
		i, err := strconv.Atoi(val)
		if err == nil {
			return i
		}
	}
	return 0
}

// AsInt64 converts value to int64.
func (v Value) AsInt64() int64 {
	switch val := v.Value.(type) {
	case int64:
		return val
	case int:
		return int64(val)
	case float64:
		return int64(val)
	case string:
		i, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			return i
		}
	}
	return 0
}

// AsUint64 converts value to uint64.
func (v Value) AsUint64() uint64 {
	switch val := v.Value.(type) {
	case uint64:
		return val
	case int:
		return uint64(val)
	case float64:
		return uint64(val)
	case string:
		u, err := strconv.ParseUint(val, 10, 64)
		if err == nil {
			return u
		}
	}
	return 0
}

// AsFloat64 converts value to float64.
func (v Value) AsFloat64() float64 {
	switch val := v.Value.(type) {
	case float64:
		return val
	case int:
		return float64(val)
	case string:
		f, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return f
		}
	}
	return 0
}

// AsBool converts value to bool.
func (v Value) AsBool() bool {
	switch val := v.Value.(type) {
	case bool:
		return val
	case string:
		b, err := strconv.ParseBool(val)
		if err == nil {
			return b
		}
	}
	return false
}

func (v Value) AsStringSlice() []string {
	switch val := v.Value.(type) {
	case []string:
		return val
	}
	return nil
}

// IsType checks if value has the same type as t.
func (v Value) IsType(t any) bool {
	return reflect.TypeOf(v.Value) == reflect.TypeOf(t)
}

func (v Value) IsArray() bool {
	return reflect.TypeOf(v.Value).Kind() == reflect.Slice
}
