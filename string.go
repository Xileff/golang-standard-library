package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("f", "felix"))
	fmt.Println(strings.Split("f e l i x", " "))
	fmt.Println(strings.ToLower("FELIX"))
	fmt.Println(strings.ToUpper("felix"))
	fmt.Println(strings.Trim("  felix  ", " "))
	fmt.Println(strings.ReplaceAll("f e l i x", " ", ""))
}
