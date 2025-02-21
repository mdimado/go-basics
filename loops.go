package main

import "fmt"

func loops() {

	//for loop
	for i := 0; i < 50; i++ {
		fmt.Println(i)
	}

	//while loop (there is no while keyword in go, for is used)
	i := 0
	for i < 5 {
		fmt.Println(i)
		i += 1
	}
}
