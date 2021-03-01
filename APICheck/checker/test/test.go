package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{Name: "Damien", Age: 30}
	v := reflect.ValueOf(user)
	f := reflect.TypeOf(user)
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Field(i))
		fmt.Println(f.Field(i).Name)
	}
}
