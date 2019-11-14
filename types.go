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

func (m *Model) validate(data []string) bool {
	t := reflect.TypeOf(data)

	for _, f := range m.fields {
		_, e := t.FieldByName(f.name)
		if e == false {
			return false
		}
	}

	return true
}

func (m *Model) getInfoFromData(data map[string]interface{}) ([]string, []interface{}) {
	keys := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))

	for k, v := range data {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}

func (m *Model) Create(data map[string]interface{}) (interface{}, bool) {
	keys, values := m.getInfoFromData(data)

	sql, _, _ := qb.Insert(m.tablename).Columns(keys...).Values(values...).ToSql()
	fmt.Println(sql)

	// execute query and get return

	return data, true
}
