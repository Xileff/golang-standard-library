package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Felix1", "Felix2", "Felix3"}
	values := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Felix1"))
}
