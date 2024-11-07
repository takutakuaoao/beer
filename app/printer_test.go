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

	assert.True(t, result.Equal(*NewContent("(struct) Empty {}")))
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

			assert.True(t, result.Equal(*NewContent(tt.expect)))
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
		PublicProperty7: func(a map[string]HasProperty) {},
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
	PublicProperty6 (string[]) [test test2]
	PublicProperty7 (func) func(map[string]HasProperty) -> void
}`

	assert.True(t, result.Equal(*NewContent(expect)))
}

func Test_to_show_func_property_of_struct(t *testing.T) {
	type HasFuncProperty struct {
		Property1 func()
	}

	s := HasFuncProperty{
		Property1: func() {},
	}

	sut := Printer{
		item: s,
	}

	result := sut.Write()

	expect := `(struct) HasFuncProperty {
	Property1 (func) func() -> void
}`

	assert.True(t, result.Equal(*NewContent(expect)))
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
	PublicProperty7 func(a map[string]HasProperty)
}
