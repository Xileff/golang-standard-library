package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type : %s", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)

		fmt.Println(valueField.Tag.Get("required"))
		fmt.Println(valueField.Tag.Get("max"))
	}
}

func isValid(value any) bool {
	result := true
	t := reflect.TypeOf(value)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""

			if result == false {
				return result
			}
		}
	}

	return true
}

func main() {
	readField(Sample{"Felix"})
	readField(Person{"Felix", "Indonesia", "felix@email.com"})
	fmt.Println("isValid :", isValid(Person{
		Name:    "Felix",
		Address: "Indonesia",
		Email:   "",
	}))
}
