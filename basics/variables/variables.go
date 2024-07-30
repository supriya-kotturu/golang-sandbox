package variables

import (
	"fmt"
	"path"
)

var PackageLevelVar int = 34

func PrintVariables() {
	var (
		name string = "John"
		age  int    = 30
	)

	dir, file := path.Split("path/to/file.txt")

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	fmt.Println("Directory:", dir)
	fmt.Println("File:", file)
}
