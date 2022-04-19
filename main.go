package main

import "fmt"

func main() {
	rules := Rules{
		&SimpleRule{
			base: 3,
			text: "Fizz",
		},
		&SimpleRule{
			base: 5,
			text: "Buzz",
		},
	}
	for i := 1; i < 15+1; i++ {
		value := ToValue(i)
		result := rules.Apply(value)
		fmt.Println(result)
	}
}
