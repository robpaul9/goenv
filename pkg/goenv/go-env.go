// Package env converts environment variable values from strings to different data types
package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetString returns a string representation of the environment variable named by the key.
//
// It returns an empty string if the variable is not present.
func GetString(key string) string {
	return os.Getenv(key)
}

// MustGetBool returns a boolean representation of the environment variable named by the key.
//
// It returns false if the variable is not present.
//
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value causes a panic.
func MustGetBool(key string) bool {
	b, err := GetBool(key)
	if err != nil {
		panic(err)
	}

	return b
}

// GetBool returns a boolean representation of the environment variable named by the key.
//
// It returns false if the variable is not present.
//
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value return false and an error.
func GetBool(key string) (bool, error) {
	e := GetString(key)

	if e == "" {
		return false, nil
	}

	i, err := strconv.ParseBool(e)
	if err != nil {
		return false, fmt.Errorf("unable to parse env key: %s with value: %v as a bool", key, e)
	}

	return i, nil
}

// MustGetInt returns an int representation of the environment variable named by the key.
//
// It returns 0 if the variable is not present.
//
// It accepts a size of 32 bits in 32 bit systems and 64 bit in 64 bit systems and
// a range from -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808
// to 9223372036854775807 in 64 bit systems all represented as strings
// Any other value causes a panic
func MustGetInt(key string) int {
	i, err := GetInt(key)
	if err != nil {
		panic(err)
	}

	return i
}

// GetInt returns an int representation of the environment variable named by the key.
//
// It returns 0 if the variable is not present.
//
// It accepts a size of 32 bits in 32 bit systems and 64 bit in 64 bit systems and
// a range from -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808
// to 9223372036854775807 in 64 bit systems all represented as strings
// Any other value returns 0 and an error.
func GetInt(key string) (int, error) {
	e := GetString(key)

	if e == "" {
		return 0, nil
	}

	i, err := strconv.Atoi(e)
	if err != nil {
		return 0, fmt.Errorf("unable to parse env key: %s with value: %v as an int", key, e)
	}

	return i, nil
}

// MustGetFloat32 returns an float32 representation of the environment variable named by the key.
//
// It returns 0.0 if the variable is not present.
//
// It accepts 32 bit floating point numbers represented as strings.
// Any other value causes a panic.
func MustGetFloat32(key string) float32 {
	f, err := GetFloat32(key)
	if err != nil {
		panic(err)
	}

	return f
}

// GetFloat32 returns an float32 representation of the environment variable named by the key.
//
// It returns 0.0 if the variable is not present.
//
// It accepts 32 bit floating point numbers represented as strings.
// Any other value returns 0.0 and an error.
func GetFloat32(key string) (float32, error) {
	e := GetString(key)

	if e == "" {
		return 0.0, nil
	}

	i, err := strconv.ParseFloat(e, 32)
	if err != nil {
		return 0.0, fmt.Errorf("unable to parse env key: %s with value: %v as a float32", key, e)
	}

	return float32(i), nil
}

// MustGetFloat64 returns an float32 representation of the environment variable named by the key.
//
// It returns 0.0 if the variable is not present.
//
// It accepts 64 bit floating point numbers represented as strings.
// Any other value causes a panic.
func MustGetFloat64(key string) float64 {
	f, err := GetFloat64(key)
	if err != nil {
		panic(err)
	}

	return f
}

// GetFloat64 returns an float32 representation of the environment variable named by the key.
//
// It returns 0.0 if the variable is not present.
//
// It accepts 64 bit floating point numbers represented as strings.
// Any other value returns 0.0 and an error
func GetFloat64(key string) (float64, error) {
	e := GetString(key)

	if e == "" {
		return 0.0, nil
	}

	i, err := strconv.ParseFloat(e, 64)
	if err != nil {
		return 0.0, fmt.Errorf("unable to parse env key: %s with value: %v as a float64", key, e)
	}

	return i, nil
}

// GetArrayOfStrings returns an array of strings generated from the environment variable named by the key.
// The array is separated by the sep argument.
// It returns nil if the variable is not present.
//
// Example:
// 	export STRING_ARR=config,hello,thing,10
// 	[config hello thing 10] = GetArrayOfStrings("STRING_ARR", ",")
func GetArrayOfStrings(key string, sep string) []string {
	e := GetString(key)

	if e == "" {
		return nil
	}

	strings := strings.Split(e, sep)

	return strings
}

// MustGetArrayOfInts returns an array of ints generated from the environment variable named by the key.
// The array is separated by the sep argument.
// It returns nil if the variable is not present.
//
// It accepts a size of 32 bits in 32 bit systems and 64 bit in 64 bit systems and
// a range from -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808
// to 9223372036854775807 in 64 bit systems all represented as strings and delimited.
// Any other value causes a panic.
//
// Example:
// 	export INT_ARR=100,4,95,103
// 	[100 4 95 103] = GetArrayOfInts("INT_ARR", ",")
func MustGetArrayOfInts(key string, sep string) []int {
	ints, err := GetArrayOfInts(key, sep)
	if err != nil {
		panic(err)
	}

	return ints
}

// GetArrayOfInts returns an array of ints generated from the environment variable named by the key.
// The array is separated by the sep argument.
// It returns nil if the variable is not present.
//
// It accepts a size of 32 bits in 32 bit systems and 64 bit in 64 bit systems and
// a range from -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808
// to 9223372036854775807 in 64 bit systems all represented as strings and delimited.
// Any other value returns nil and an error
//
// Example:
// 	export INT_ARR=100,4,95,103
// 	[100 4 95 103] = GetArrayOfInts("INT_ARR", ",")
func GetArrayOfInts(key string, sep string) ([]int, error) {
	var ints []int

	s := GetArrayOfStrings(key, sep)

	if s == nil {
		return nil, nil
	}

	for n, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("unable to parse env key: %s with value: %v as an int", key, s[n])
		}

		ints = append(ints, i)
	}

	return ints, nil
}

// MustGetArrayOfFloat32s returns an array of float32s generated from the environment variable named by the key.
// The array is separated by the sep argument.
// It returns nil if the variable is not present.
//
// It accepts 32 bit floating point numbers represented as strings.
// Any other value causes a panic.
//
// Example:
// 	export FLOAT32_ARR=100.4,4,95.3
// 	[100.4 4 95.3] = GetArrayOfFloat32s("FLOAT32_ARR", ",")
func MustGetArrayOfFloat32s(key string, sep string) []float32 {
	floats, err := GetArrayOfFloat32s(key, sep)
	if err != nil {
		panic(err)
	}

	return floats
}

// GetArrayOfFloat32s returns an array of float32s generated from the environment variable named by the key.
// The array is separated by the sep argument.
// It returns nil if the variable is not present.
//
// It accepts 32 bit floating point numbers represented as strings.
// Any other value return nil and an error
//
// Example:
// 	export FLOAT32_ARR=100.4,4,95.3
// 	[100.4 4 95.3] = GetArrayOfFloat32s("FLOAT32_ARR", ",")
func GetArrayOfFloat32s(key string, sep string) ([]float32, error) {
	var floats []float32

	s := GetArrayOfStrings(key, sep)

	if s == nil {
		return nil, nil
	}

	for n, v := range s {
		i, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil, fmt.Errorf("unable to parse env key: %s with value: %v as an float32", key, s[n])
		}

		floats = append(floats, float32(i))
	}

	return floats, nil
}

// MustGetArrayOfFloat64s returns an array of float32s generated from the environment variable named by the key.
// The array is separated by the sep argument.
// It returns nil if the variable is not present.
//
// It accepts 64 bit floating point numbers represented as strings.
// Any other value causes a panic.
//
// Example:
// 	export FLOAT64_ARR=100.4,4,95.3
// 	[100.4 4 95.3] = GetArrayOfFloat64s("FLOAT64_ARR", ",")
func MustGetArrayOfFloat64s(key string, sep string) []float64 {
	floats, err := GetArrayOfFloat64s(key, sep)
	if err != nil {
		panic(err)
	}

	return floats
}

// GetArrayOfFloat64s returns an array of float32s generated from the environment variable named by the key.
// The array is separated by the sep argument.
// It returns nil if the variable is not present.
//
// It accepts 64 bit floating point numbers represented as strings.
// Any other value causes a panic.
//
// Example:
// 	export FLOAT64_ARR=100.4,4,95.3
// 	[100.4 4 95.3] = GetArrayOfFloat64s("FLOAT64_ARR", ",")
func GetArrayOfFloat64s(key string, sep string) ([]float64, error) {
	var floats []float64

	s := GetArrayOfStrings(key, sep)

	if s == nil {
		return nil, nil
	}

	for n, v := range s {
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse env key: %s with value: %v as an float64", key, s[n])
		}

		floats = append(floats, i)
	}

	return floats, nil
}
