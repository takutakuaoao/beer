package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Property interface {
	GetText() string
}

func NewProperty(name string, value reflect.Value, kind reflect.Type) Property {
	if kind.Kind() == reflect.Slice {
		return NewSliceProperty(name, kind, value)
	}

	if kind.Kind() == reflect.Func {
		return NewFuncProperty(name, kind)
	}

	return NewSimplyProperty(name, kind, value)
}

type FuncProperty struct {
	name string
	kind reflect.Type
}

func NewFuncProperty(name string, kind reflect.Type) FuncProperty {
	return FuncProperty{
		name: name,
		kind: kind,
	}
}

func (p FuncProperty) GetText() string {
	funcArgType := ""
	returnType := "void"

	if p.hasArg() {
		funcArgType = p.getAllArgTypes()
	}

	if p.hasReturnType() {
		returnType = p.getAllReturnTypes()
	}

	return formatProperty(p.name, "func", fmt.Sprintf("func(%s) -> %s", funcArgType, returnType))
}

func (p *FuncProperty) hasArg() bool {
	return p.kind.NumIn() > 0
}

func (p *FuncProperty) getAllArgTypes() string {
	funcArgType := ""

	for i := 0; i < p.kind.NumIn(); i++ {
		typeText := getTypeText(p.kind.In(i))

		if i == 0 {
			funcArgType = typeText
		} else {
			funcArgType = fmt.Sprintf("%s, %s", funcArgType, typeText)
		}
	}

	return funcArgType
}

func (p *FuncProperty) hasReturnType() bool {
	return p.kind.NumOut() > 0
}

func (p *FuncProperty) getAllReturnTypes() string {
	typeTexts := []string{}

	for i := 0; i < p.kind.NumOut(); i++ {
		typeTexts = append(typeTexts, getTypeText(p.kind.Out(i)))
	}

	typeText := strings.Join(typeTexts, ", ")

	if len(typeTexts) == 1 {
		return typeText
	}

	return fmt.Sprintf("(%s)", typeText)
}

func getTypeText(t reflect.Type) string {
	if isSliceType(t) {
		return formatSliceTypeText(t.Elem())
	}

	if isMapType(t) {
		return formatMapTypeText(t.Key(), t.Elem().Name())
	}

	return t.Name()
}

func isSliceType(t reflect.Type) bool {
	return t.Kind() == reflect.Slice
}

func isMapType(t reflect.Type) bool {
	return t.Kind() == reflect.Map
}

type SliceProperty struct {
	name  string
	kind  reflect.Type
	value reflect.Value
}

func NewSliceProperty(name string, kind reflect.Type, value reflect.Value) SliceProperty {
	return SliceProperty{
		name:  name,
		kind:  kind,
		value: value,
	}
}

func (p SliceProperty) GetText() string {
	return formatProperty(p.name, formatSliceTypeText(p.kind.Elem()), fmt.Sprintf("%v", p.value))
}

type SimplyProperty struct {
	name  string
	kind  string
	value string
}

func NewSimplyProperty(name string, kind reflect.Type, value reflect.Value) SimplyProperty {
	return SimplyProperty{
		name:  name,
		kind:  kind.Name(),
		value: fmt.Sprintf("%v", value),
	}
}

func (p SimplyProperty) GetText() string {
	value := ""

	if p.kind == "string" {
		value = fmt.Sprintf("\"%s\"", p.value)
	} else {
		value = p.value
	}

	return formatProperty(p.name, p.kind, value)
}

func formatProperty(name string, kind string, value string) string {
	return fmt.Sprintf("%s (%s) %s", name, kind, value)
}

func formatSliceTypeText(sliceType reflect.Type) string {
	return fmt.Sprintf("%s[]", sliceType.Name())
}

func formatMapTypeText(keyType reflect.Type, valueType string) string {
	return fmt.Sprintf("map[%s]%s", keyType, valueType)
}
