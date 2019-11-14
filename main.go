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

	response, _ := user.Create(User{
		nome:  "Matheus Henrique dos Santos",
		email: "matherique@gmail.com",
		senha: "senhateste",
	})

	fmt.Println(response)
}
