package main

import (
	"errors"

	slic6 "github.com/ashcoder666/boot.dev.go/contents/slice-chapter/sli-c6"
)

// loop4 "github.com/ashcoder666/boot.dev.go/contents/loops-chapter/loop-4"
// topic1 "github.com/ashcoder666/boot.dev.go/contents/structs-3"
// structs3 "github.com/ashcoder666/boot.dev.go/contents/structs-chapter/structs-3"

func main() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	slic6.Test("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	slic6.Test("Error on mail server", mailErrors, commaDelimit)
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}
