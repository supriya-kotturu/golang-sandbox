package main

import "fmt"

func checkEvenOrOdd(list []int) {
	for _, ele := range list {
		if ele%2 == 0 {
			fmt.Println(ele, " is Even")
		} else {
			fmt.Println(ele, " is Odd")
		}
	}
}

func main() {
	num := []int{3, 89, 6, 7, 3, 5, 8, 2}
	checkEvenOrOdd(num)
}
