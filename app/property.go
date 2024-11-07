package main

import (
	"fmt"
	"reflect"
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
	fv := ""

	if p.kind.NumIn() > 0 {
		arg := p.kind.In(0)

		if arg.Kind() == reflect.Slice {
			fv = formatSliceTypeText(arg.Elem())
		} else {
			fv = arg.Name()
		}
	}

	return formatProperty(p.name, "func", fmt.Sprintf("func(%s) -> void", fv))
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

func formatSliceTypeText(slice reflect.Type) string {
	return fmt.Sprintf("%s[]", slice)
}
