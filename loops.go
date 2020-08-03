package main

import "fmt"

func main() {
	var i int

	for i <= 100 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 100; j >= 0; j-- {
		fmt.Println(j)
	}

	k := 4

	for {
		println(k)
		if k == 0 {
			break
		}
		k--
	}
}
