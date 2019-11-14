package main

import (
	"fmt"
)

// Tipo inteiro
type Int struct {
	defaultValue int
}

// Tipo string
type String struct {
	Size         int
	defaultValue string
}

// Tipo Data
type Date struct {
	defaultValue string
}

// Tipo Data Only
type DateOnly struct {
	defaultValue string
}

type Field struct {
	name string
	t    interface{}
}

type Model struct {
	fields []Field
}

func (m *Model) Create(data interface{}) interface{} {
	for _, f := range m.fields {
		fmt.Println(f.name)
	}
	fmt.Println("cadastrar algo")
	return data
}
