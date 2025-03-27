package main

import (
	"errors"
	"fmt"
)

// loop4 "github.com/ashcoder666/boot.dev.go/contents/loops-chapter/loop-4"
// topic1 "github.com/ashcoder666/boot.dev.go/contents/structs-3"
// structs3 "github.com/ashcoder666/boot.dev.go/contents/structs-chapter/structs-3"

func main() {

	// map study
	fmt.Println(getUserMap([]string{"Eren", "Armin", "Mikasa"}, []int{14355550987, 98765550987, 18265554567}))
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	detailMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	for _, name := range names {
		for _, number := range phoneNumbers {
			detailMap[name] = user{
				name:        name,
				phoneNumber: number,
			}
		}
	}

	return detailMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
