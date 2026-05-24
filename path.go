package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Dir("test/world.go"))
	fmt.Println(path.Base("test/world.go"))
	fmt.Println(path.Ext("test/world.go"))
	fmt.Println(filepath.IsAbs("test/world.go"))
	fmt.Println(filepath.IsLocal("test/world.go"))
	fmt.Println(path.Join("test1", "test2", "test3"))
}
