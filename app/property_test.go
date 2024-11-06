package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
