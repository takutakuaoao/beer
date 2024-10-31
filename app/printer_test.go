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

func Test_to_show_multiple_properties_of_struct(t *testing.T) {
	object := HasMultipleProperties{
		PublicProperty1: "test",
		PublicProperty2: 1,
		PublicProperty3: true,
		PublicProperty4: -1,
		PublicProperty5: 10,
		PublicProperty6: []string{"test", "test2"},
	}

	sut := Printer{
		item: object,
	}

	result := sut.Write()

	expect := `(struct) HasMultipleProperties {
	PublicProperty1 (string) "test"
	PublicProperty2 (uint8) 1
	PublicProperty3 (bool) true
	PublicProperty4 (int) -1
	PublicProperty5 (uint8) 10
	PublicProperty6 (string[]) ["test", "test2"]
}`

	assert.Equal(t, expect, result)
}

type Empty struct {
}

type HasProperty struct {
	PublicProperty string
}

type HasMultipleProperties struct {
	PublicProperty1 string
	PublicProperty2 uint8
	PublicProperty3 bool
	PublicProperty4 int
	PublicProperty5 byte
	PublicProperty6 []string
}
