package structs3

import "fmt"

func Examplea() {

	myCar := struct {
		brand string
		model string
	}{
		brand: "Toyota",
		model: "ss",
	}

	fmt.Println(myCar)
}
