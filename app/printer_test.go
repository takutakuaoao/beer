package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_to_show_empty_struct(t *testing.T) {
	sut := Printer{
		item: Empty{},
	}

	result := sut.Write()

	assert.Equal(t, "(struct) Empty {}", result)
}

type Empty struct {
}
