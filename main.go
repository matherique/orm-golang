package main

import (
	"fmt"
)

type User struct {
	id    int    ``
	nome  string ``
	email string ``
	senha string ``
}

var user = Model{fields: []Field{
	{
		name: "id",
		t:    Int{},
	},
	{
		name: "nome",
		t:    String{Size: 64},
	},
	{
		name: "email",
		t:    String{Size: 32},
	},
	{
		name: "senha",
		t:    String{Size: 32},
	},
},
	tablename: "user",
}

func main() {
	var userdata = make(map[string]interface{})

	userdata["nome"] = "Matheus Henrique dos Santos"
	userdata["email"] = "matherique@gmail.com"
	userdata["senha"] = "senhateste"

	fmt.Println(userdata)

	resp, _ := user.Create(userdata)
	fmt.Println(resp)
}
