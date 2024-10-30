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

func Test_to_show_properties_of_struct(t *testing.T) {
	cases := []struct {
		name   string
		item   interface{}
		expect string
	}{
		{
			name: "with value",
			item: HasProperty{
				PublicProperty: "test",
			},
			expect: "(struct) HasProperty {\n\tPublicProperty (string) \"test\"\n}",
		},
		{
			name:   "with non value",
			item:   HasProperty{},
			expect: "(struct) HasProperty {\n\tPublicProperty (string) \"\"\n}",
		},
		{
			name: "with multiple properties",
			item: HasMultipleProperties{
				PublicProperty2: 4,
			},
			expect: "(struct) HasMultipleProperties {\n\tPublicProperty1 (string) \"\"\n\tPublicProperty2 (uint8) 4\n}",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			sut := Printer{
				item: tt.item,
			}

			result := sut.Write()

			assert.Equal(t, tt.expect, result)
		})
	}
}

type Empty struct {
}

type HasProperty struct {
	PublicProperty string
}

type HasMultipleProperties struct {
	PublicProperty1 string
	PublicProperty2 uint8
}
