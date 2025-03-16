package loop4

import (
	"fmt"
)

func Fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz", i)
		} else if i%5 == 0 {
			fmt.Println("buzz", i)
		} else if i%3 == 0 {
			fmt.Println("fizz", i)
		}
	}
}

// func printPrimes(max int) {
// 	fmt.Println(1)
// 	fmt.Println(2)
// 	for i:=3 ; i< math.Sqrt2(n)
// }

// // don't edit below this line

// func Test(max int) {
// 	fmt.Printf("Primes up to %v:\n", max)
// 	printPrimes(max)
// 	fmt.Println("===============================================================")
// }

// func main() {
// 	test(10)
// 	test(20)
// 	test(30)
// }
