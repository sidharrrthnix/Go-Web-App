package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "John Doe",
		Bio:  "I am a software engineer",
		Age:  30,
	}

	if err := t.Execute(os.Stdout, user); err != nil {
		panic(err)
	}

}
