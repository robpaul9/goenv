package env_test

import (
	"os"
	"testing"

	env "github.com/robpaul9/goenv"
	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	os.Setenv("STRING_ENV", "hello")
	defer os.Unsetenv("STRING_ENV")

	computedValue := env.GetString("STRING_ENV")
	expectedValue := "hello"

	assert.Equal(t, expectedValue, computedValue)
}

func TestMustGetBool(t *testing.T) {
	os.Setenv("BOOL_ENV", "true")
	defer os.Unsetenv("BOOL_ENV")

	computedValue := env.MustGetBool("BOOL_ENV")

	assert.True(t, computedValue)
}

func TestMustGetBoolEmptyEnvVar(t *testing.T) {
	computedValue := env.MustGetBool("BOOL_ENV")

	assert.False(t, computedValue)
}

func TestMustGetBoolError(t *testing.T) {
	os.Setenv("BOOL_ENV", "string-type")
	defer os.Unsetenv("BOOL_ENV")

	assert.Panics(t, func() {
		env.MustGetBool("BOOL_ENV")
	})
}

func TestGetBool(t *testing.T) {
	os.Setenv("BOOL_ENV", "true")
	defer os.Unsetenv("BOOL_ENV")

	computedValue, _ := env.GetBool("BOOL_ENV")

	assert.True(t, computedValue)
}

func TestGetBoolEmptyEnvVar(t *testing.T) {
	computedValue, _ := env.GetBool("BOOL_ENV")

	assert.False(t, computedValue)
}

func TestGetBoolError(t *testing.T) {
	os.Setenv("BOOL_ENV", "string-type")
	defer os.Unsetenv("BOOL_ENV")

	_, err := env.GetBool("BOOL_ENV")

	assert.NotNil(t, err)
}

func TestMustGetInt(t *testing.T) {
	os.Setenv("INT_ENV", "100")
	defer os.Unsetenv("INT_ENV")

	computedValue := env.MustGetInt("INT_ENV")

	assert.Equal(t, 100, computedValue)
}

func TestMustGetIntEmptyEnvVar(t *testing.T) {
	computedValue := env.MustGetInt("BOOL_ENV")

	assert.Equal(t, 0, computedValue)
}

func TestMustGetIntError(t *testing.T) {
	os.Setenv("INT_ENV", "string-type")
	defer os.Unsetenv("INT_ENV")

	assert.Panics(t, func() {
		env.MustGetInt("INT_ENV")
	})
}

func TestGetInt(t *testing.T) {
	os.Setenv("INT_ENV", "100")
	defer os.Unsetenv("INT_ENV")

	computedValue, _ := env.GetInt("INT_ENV")

	assert.Equal(t, 100, computedValue)
}

func TestGetIntEmptyEnvVar(t *testing.T) {
	computedValue, _ := env.GetInt("BOOL_ENV")

	assert.Equal(t, 0, computedValue)
}

func TestGetIntError(t *testing.T) {
	os.Setenv("INT_ENV", "string-type")
	defer os.Unsetenv("INT_ENV")

	_, err := env.GetInt("INT_ENV")

	assert.NotNil(t, err)
}

func TestMustGetFloat32(t *testing.T) {
	os.Setenv("FLOAT32_ENV", "100.100")
	defer os.Unsetenv("FLOAT32_ENV")

	computedValue := env.MustGetFloat32("FLOAT32_ENV")

	assert.Equal(t, float32(100.100), computedValue)
}

func TestMustGetFloat32EmptyEnvVar(t *testing.T) {
	computedValue := env.MustGetFloat32("FLOAT32_ENV")

	assert.Equal(t, float32(0.0), computedValue)
}

func TestMustGetFloat32Error(t *testing.T) {
	os.Setenv("FLOAT32_ENV", "string-type")
	defer os.Unsetenv("FLOAT32_ENV")

	assert.Panics(t, func() {
		env.MustGetFloat32("FLOAT32_ENV")
	})
}

func TestGetFloat32(t *testing.T) {
	os.Setenv("FLOAT32_ENV", "100.100")
	defer os.Unsetenv("FLOAT32_ENV")

	computedValue, _ := env.GetFloat32("FLOAT32_ENV")

	assert.Equal(t, float32(100.100), computedValue)
}

func TestGetFloat32EmptyEnvVar(t *testing.T) {
	computedValue, _ := env.GetFloat32("FLOAT32_ENV")

	assert.Equal(t, float32(0.0), computedValue)
}

func TestGetFloat32Error(t *testing.T) {
	os.Setenv("FLOAT32_ENV", "string-type")
	defer os.Unsetenv("FLOAT32_ENV")

	_, err := env.GetFloat32("FLOAT32_ENV")

	assert.NotNil(t, err)
}

func TestMustGetFloat64(t *testing.T) {
	os.Setenv("FLOAT64_ENV", "100.100")
	defer os.Unsetenv("FLOAT64_ENV")

	computedValue := env.MustGetFloat64("FLOAT64_ENV")

	assert.Equal(t, 100.100, computedValue)
}

func TestMustGetFloat64EmptyEnvVar(t *testing.T) {
	computedValue := env.MustGetFloat64("FLOAT64_ENV")

	assert.Equal(t, 0.0, computedValue)
}

func TestMustGetFloat64Error(t *testing.T) {
	os.Setenv("FLOAT64_ENV", "string-type")
	defer os.Unsetenv("FLOAT64_ENV")

	assert.Panics(t, func() {
		env.MustGetFloat64("FLOAT64_ENV")
	})
}

func TestGetFloat64(t *testing.T) {
	os.Setenv("FLOAT64_ENV", "100.100")
	defer os.Unsetenv("FLOAT64_ENV")

	computedValue, _ := env.GetFloat64("FLOAT64_ENV")

	assert.Equal(t, 100.100, computedValue)
}

func TestGetFloat64EmptyEnvVar(t *testing.T) {
	computedValue, _ := env.GetFloat64("FLOAT64_ENV")

	assert.Equal(t, 0.0, computedValue)
}

func TestGetFloat64Error(t *testing.T) {
	os.Setenv("FLOAT64_ENV", "string-type")
	defer os.Unsetenv("FLOAT64_ENV")

	_, err := env.GetFloat64("FLOAT64_ENV")

	assert.NotNil(t, err)
}

func TestGetArrayOfStrings(t *testing.T) {
	os.Setenv("ARR_STRINGS", "hello,from,go")
	defer os.Unsetenv("ARR_STRINGS")

	expectedValue := []string{"hello", "from", "go"}

	computedValue := env.GetArrayOfStrings("ARR_STRINGS", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestGetArrayOfStringsEmptyEnvVar(t *testing.T) {
	var expectedValue []string

	computedValue := env.GetArrayOfStrings("ARR_STRINGS", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestMustGetArrayOfInts(t *testing.T) {
	os.Setenv("ARR_INTS", "10,3,5")
	defer os.Unsetenv("ARR_INTS")

	expectedValue := []int{10, 3, 5}

	computedValue := env.MustGetArrayOfInts("ARR_INTS", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestMustGetArrayOfIntsEmptyEnvVar(t *testing.T) {
	var expectedValue []int

	computedValue := env.MustGetArrayOfInts("ARR_INTS", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestMustGetArrayOfIntsError(t *testing.T) {
	os.Setenv("ARR_INTS", "10,3,5.5")
	defer os.Unsetenv("ARR_INTS")

	assert.Panics(t, func() {
		env.MustGetArrayOfInts("ARR_INTS", ",")
	})
}

func TestGetArrayOfInts(t *testing.T) {
	os.Setenv("ARR_INTS", "10,3,5")
	defer os.Unsetenv("ARR_INTS")

	expectedValue := []int{10, 3, 5}

	computedValue, _ := env.GetArrayOfInts("ARR_INTS", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestGetArrayOfIntsEmptyEnvVar(t *testing.T) {
	var expectedValue []int

	computedValue, _ := env.GetArrayOfInts("ARR_INTS", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestGetArrayOfIntsError(t *testing.T) {
	os.Setenv("ARR_INTS", "10,3,5.5")
	defer os.Unsetenv("ARR_INTS")

	_, err := env.GetArrayOfInts("ARR_INTS", ",")

	assert.NotNil(t, err)
}

func TestMustGetArrayOfFloat32s(t *testing.T) {
	os.Setenv("ARR_FLOAT32s", "10.10,3,5.40101")
	defer os.Unsetenv("ARR_FLOAT32s")

	expectedValue := []float32{10.10, 3, 5.40101}

	computedValue := env.MustGetArrayOfFloat32s("ARR_FLOAT32s", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestMustGetArrayOfFloat32sEmptyEnvVar(t *testing.T) {
	var expectedValue []float32

	computedValue := env.MustGetArrayOfFloat32s("ARR_FLOAT32s", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestMustGetArrayOfFloat32sError(t *testing.T) {
	os.Setenv("ARR_FLOAT32s", "10,3,sneaky-string,100")
	defer os.Unsetenv("ARR_FLOAT32s")

	assert.Panics(t, func() {
		env.MustGetArrayOfFloat32s("ARR_FLOAT32s", ",")
	})
}

func TestGetArrayOfFloat32s(t *testing.T) {
	os.Setenv("ARR_FLOAT32s", "10.10,3,5.40101")
	defer os.Unsetenv("ARR_FLOAT32s")

	expectedValue := []float32{10.10, 3, 5.40101}

	computedValue, _ := env.GetArrayOfFloat32s("ARR_FLOAT32s", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestGetArrayOfFloat32sEmptyEnvVar(t *testing.T) {
	var expectedValue []float32

	computedValue, _ := env.GetArrayOfFloat32s("ARR_FLOAT32s", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestGetArrayOfFloat32sError(t *testing.T) {
	os.Setenv("ARR_FLOAT32s", "10,3,sneaky-string,100")
	defer os.Unsetenv("ARR_FLOAT32s")

	_, err := env.GetArrayOfFloat32s("ARR_FLOAT32s", ",")

	assert.NotNil(t, err)
}

func TestMustGetArrayOfFloat64s(t *testing.T) {
	os.Setenv("ARR_FLOAT64s", "10.10,3,5.40101")
	defer os.Unsetenv("ARR_FLOAT64s")

	expectedValue := []float64{10.10, 3, 5.40101}

	computedValue := env.MustGetArrayOfFloat64s("ARR_FLOAT64s", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestMustGetArrayOfFloat64sEmptyEnvVar(t *testing.T) {
	var expectedValue []float64

	computedValue := env.MustGetArrayOfFloat64s("ARR_FLOAT64s", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestMustGetArrayOfFloat64sError(t *testing.T) {
	os.Setenv("ARR_FLOAT64s", "10,3,sneaky-string,100")
	defer os.Unsetenv("ARR_FLOAT64s")

	assert.Panics(t, func() {
		env.MustGetArrayOfFloat64s("ARR_FLOAT64s", ",")
	})
}

func TestGetArrayOfFloat64s(t *testing.T) {
	os.Setenv("ARR_FLOAT64s", "10.10,3,5.40101")
	defer os.Unsetenv("ARR_FLOAT64s")

	expectedValue := []float64{10.10, 3, 5.40101}

	computedValue, _ := env.GetArrayOfFloat64s("ARR_FLOAT64s", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestGetArrayOfFloat64sEmptyEnvVar(t *testing.T) {
	var expectedValue []float64

	computedValue, _ := env.GetArrayOfFloat64s("ARR_FLOAT64s", ",")

	assert.Equal(t, expectedValue, computedValue)
}

func TestGetArrayOfFloat64sError(t *testing.T) {
	os.Setenv("ARR_FLOAT64s", "10,3,sneaky-string,100")
	defer os.Unsetenv("ARR_FLOAT64s")

	_, err := env.GetArrayOfFloat64s("ARR_FLOAT64s", ",")

	assert.NotNil(t, err)
}
