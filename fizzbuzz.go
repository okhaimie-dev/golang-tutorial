//fizz buzz program in go

package main

import(
	"fmt"
)

func main(){
	for i := 1; i < 21; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizz buzz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}
