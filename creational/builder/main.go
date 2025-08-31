package main

import "fmt"

type Object struct {
	name string
	data string
}

func (o *Object) Print() {
	fmt.Printf("Name: %s; Data: %s;\n", o.name, o.data)
}

type ObjectBuilder struct {
	name string
	data string
}

func NewBuilder() *ObjectBuilder {
	return &ObjectBuilder{
		name: "default_name",
		data: "default_data",
	}
}

func (b *ObjectBuilder) WithName(name string) *ObjectBuilder {
	b.name = name
	return b
}

func (b *ObjectBuilder) WithData(data string) *ObjectBuilder {
	b.data = data
	return b
}

func (b *ObjectBuilder) Build() *Object {
	return &Object{
		name: b.name,
		data: b.data,
	}
}

func main() {
	obj := NewBuilder().Build()
	obj.Print()

	obj = NewBuilder().
		WithData("data").
		WithName("name").
		Build()

	obj.Print()
}
