package main

import (
	"fmt"
	"reflect"
)

type Printer struct {
	item interface{}
}

func (p *Printer) Write() string {
	o := reflect.TypeOf(p.item)

	return fmt.Sprintf("(%v) %v {}", o.Kind(), o.Name())
}
