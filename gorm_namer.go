package main

import (
	"bytes"
	"fmt"
)

func main() {

	fmt.Println("CategoryName:", leeNamer("CategoryName"))

	fmt.Println("Category1name:", leeNamer("Category1name"))
	fmt.Println("category_name:", leeNamer("category_name"))

	fmt.Println("_category_name_:", leeNamer("_category_name_"))

}

func leeNamer(name string) string {
	const (
		lower = false
		upper = true
	)

	var (
		value    = name
		buf      = bytes.NewBufferString("")
		currCase bool
	)

	for i, v := range value {
		currCase = bool(value[i] >= 'A' && value[i] <= 'Z')
		if i == 0 && currCase == upper {
			buf.WriteRune(v + 32)
		} else {
			buf.WriteRune(v)
		}
	}

	return buf.String()
}
