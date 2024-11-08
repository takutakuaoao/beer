package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_to_get_text_of_the_Func_property(t *testing.T) {
	cases := []struct {
		name     string
		f        interface{}
		expected string
	}{
		{
			name:     "one arg",
			f:        func(_ string) {},
			expected: "Property (func) func(string) -> void",
		},
		{
			name:     "one arg with slice",
			f:        func(_ []string) {},
			expected: "Property (func) func(string[]) -> void",
		},
		{
			name:     "one arg with map",
			f:        func(_ map[string]uint8) {},
			expected: "Property (func) func(map[string]uint8) -> void",
		},
		{
			name:     "one arg with struct",
			f:        func(_ SampleStruct) {},
			expected: "Property (func) func(SampleStruct) -> void",
		},
		{
			name:     "multiple args",
			f:        func(_ map[string]SampleStruct, _ []SampleStruct) {},
			expected: "Property (func) func(map[string]SampleStruct, SampleStruct[]) -> void",
		},
		{
			name:     "return one value",
			f:        func() string { return "" },
			expected: "Property (func) func() -> string",
		},
		{
			name:     "return one value with map",
			f:        func() map[string]SampleStruct { return map[string]SampleStruct{} },
			expected: "Property (func) func() -> map[string]SampleStruct",
		},
		{
			name:     "return one value with Struct",
			f:        func() SampleStruct { return SampleStruct{} },
			expected: "Property (func) func() -> SampleStruct",
		},
		{
			name:     "return one value with slice",
			f:        func() []SampleStruct { return []SampleStruct{} },
			expected: "Property (func) func() -> SampleStruct[]",
		},
		{
			name:     "return multiple types",
			f:        func() (string, error) { return "", nil },
			expected: "Property (func) func() -> (string, error)",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			sut := NewFuncProperty("Property", reflect.TypeOf(tt.f))

			assert.Equal(t, tt.expected, sut.GetText())
		})
	}
}

// This is the Learning Test.
func Test_to_make_the_reflection_type_from_func_variable(t *testing.T) {
	f := func(test string) string {
		return ""
	}

	tf := reflect.TypeOf(f)

	// get type
	assert.Equal(t, "func", tf.Kind().String())

	// get count args
	assert.Equal(t, 1, tf.NumIn())

	// get count returns
	assert.Equal(t, 1, tf.NumOut())

	// get arg type
	assert.Equal(t, "string", tf.In(0).String())

	// get return type
	assert.Equal(t, "string", tf.Out(0).String())
}

type SampleStruct struct {
	Property string
}
