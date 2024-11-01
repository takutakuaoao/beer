package main

import (
	"fmt"
	"reflect"
)

type Printer struct {
	item interface{}
}

func (p *Printer) Write() string {
	object := NewStruct(reflect.TypeOf(p.item), reflect.ValueOf(p.item))
	content := *NewContent("")

	if object.HasProperty() {
		content = object.WriteProperties(content)
	}

	return object.WriteStruct(content).content
}

func NewStruct(reflectType reflect.Type, reflectValue reflect.Value) *Struct {
	return &Struct{
		reflectType:  reflectType,
		reflectValue: reflectValue,
	}
}

type Struct struct {
	reflectType  reflect.Type
	reflectValue reflect.Value
}

func (s *Struct) HasProperty() bool {
	return s.reflectType.NumField() > 0
}

func (s *Struct) WriteStruct(content Content) Content {
	identified := fmt.Sprintf("(%v) %v", s.getType(), s.getStructName())

	return content.Surround("{", "}").AddHead(fmt.Sprintf("%v ", identified))
}

func (s *Struct) getStructName() string {
	return s.reflectType.Name()
}

func (s *Struct) getType() string {
	return s.reflectType.Kind().String()
}

func (s *Struct) WriteProperties(content Content) Content {
	s.loopProperties(func(name string, propertyType *PropertyType, propertyValue *PropertyValue) {
		content = content.Indent().Merge(
			s.propertyText(name, propertyType, propertyValue),
		)
	})

	return content.LineBreak()
}

func (s *Struct) loopProperties(callBack func(name string, propertyType *PropertyType, propertyValue *PropertyValue)) {
	for i := 0; i < s.reflectType.NumField(); i++ {
		field := s.reflectType.Field(i)

		callBack(
			field.Name,
			NewPropertyType(field.Type),
			NewPropertyValue(s.reflectValue.FieldByName(field.Name)),
		)
	}
}

func (s *Struct) propertyText(name string, propertyType *PropertyType, propertyValue *PropertyValue) string {
	return fmt.Sprintf("%v (%v) %v", name, propertyType, propertyValue.GetAsText(propertyType))
}

type PropertyType struct {
	self reflect.Type
}

func NewPropertyType(value reflect.Type) *PropertyType {
	return &PropertyType{
		self: value,
	}
}

func (t *PropertyType) IsString() bool {
	return t.self.Kind() == reflect.String
}

func (t *PropertyType) String() string {
	return t.getTypeName()
}

func (t *PropertyType) getTypeName() string {
	if t.self.Kind() == reflect.Slice {
		return fmt.Sprintf("%v[]", t.self.Elem())
	}

	return t.self.Name()
}

type PropertyValue struct {
	value reflect.Value
}

func NewPropertyValue(value reflect.Value) *PropertyValue {
	return &PropertyValue{
		value: value,
	}
}

func (v *PropertyValue) GetAsText(propertyType *PropertyType) string {
	if propertyType.IsString() {
		return fmt.Sprintf("\"%v\"", v.value)
	}

	return fmt.Sprintf("%v", v.value)
}

func NewContent(content string) *Content {
	return &Content{
		content: content,
	}
}

type Content struct {
	content string
}

func (c Content) Merge(merged string) Content {
	result := fmt.Sprintf("%v%v", c.content, merged)

	return *NewContent(result)
}

func (c Content) Indent() Content {
	result := fmt.Sprintf("%v\n\t", c.content)

	return *NewContent(result)
}

func (c Content) LineBreak() Content {
	return c.Merge("\n")
}

func (c Content) Surround(pre string, last string) Content {
	return *NewContent(fmt.Sprintf("%v%v%v", pre, c.content, last))
}

func (c Content) AddHead(text string) Content {
	return *NewContent(fmt.Sprintf("%v%v", text, c.content))
}
