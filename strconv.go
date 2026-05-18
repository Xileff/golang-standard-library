package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("Error :", err.Error())
	}

	resultInt, errInt := strconv.Atoi("1000")
	if errInt == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println("Error :", errInt.Error())
	}

	result = false
	fmt.Println(result)
}
