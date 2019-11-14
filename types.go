package main

import (
	"fmt"
	qb "github.com/Masterminds/squirrel"
	"reflect"
	//"strings"
)

type Int struct{ defaultValue int }
type String struct {
	Size         int
	defaultValue string
}
type Field struct {
	name string
	t    interface{}
}
type Model struct {
	fields    []Field
	tablename string
}

func (m *Model) validate(data interface{}) bool {
	t := reflect.TypeOf(data)

	for _, f := range m.fields {
		_, e := t.FieldByName(f.name)
		if e == false {
			return false
		}
	}

	return true
}

func (m *Model) getFields(data interface{}) []string {
	var fields []string
	t := reflect.TypeOf(data)

	for _, f := range m.fields {
		fi, _ := t.FieldByName(f.name)
		fields = append(fields, fi.Name)
	}

	return fields
}

func (m *Model) Create(data interface{}) (interface{}, bool) {
	v := m.validate(data)
	if v == false {
		return nil, false
	}

	fields := m.getFields(data)

	inputsql := []string{"insert", "into", m.tablename}
	inputsql = append(inputsql, fields...)

  sql, _, _ := qb.Insert(m.tablename).Columns("aa").ToSql()
	fmt.Println(inputsql)
  fmt.Println(sql)
	return data, true
}
