package main

import "fmt"

func fizzIf() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz", i)
		} else if i%3 == 0 {
			fmt.Println("Fizz", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz", i)
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzIf()

	if num := 7; num == 7 {
		fmt.Println("yeah the damn number is ", 7)
	}

	var k bool
	switch k {
	case true:
		fmt.Println("it's f*cking true ☹🙃😋")
	case false:
		fmt.Println("it's a f*cking lie 😏😞😒")
	default:
		fmt.Println("mehn, i f*cking don't know 🤧🤐")
	}
}
