package main

import (
	"fmt"
	"strconv"
)

var userInput int
var message string

func toOneThousand() {
	numberCount := 0
	for i := 0; i <= 1000; i++ {
		fmt.Println(i)
		numberCount++
	}
	message = "There are " + strconv.Itoa(numberCount) + " numbers"
}

func oddToOneThousand() {
	numberCount := 0
	for i := 0; i <= 1000; i++ {
		if i%2 == 1 {
			fmt.Println(i)
			numberCount++
		} else {
			continue
		}
	}
	message = "There are " + strconv.Itoa(numberCount) + " numbers"
}

func evenToOneThousand() {
	numberCount := 0
	for i := 0; i <= 1000; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			numberCount++
		} else {
			continue
		}
	}
	message = "There are " + strconv.Itoa(numberCount) + " numbers"
}

func main() {
	fmt.Println("Please choose the desired number count")
	fmt.Println("1 : count to 1000 \n2 : odd numbers only count to 100 \n3 : even numbers only count to 1000")
	fmt.Scan(&userInput)

	if userInput == 1 {
		toOneThousand()
	} else if userInput == 2 {
		oddToOneThousand()
	} else if userInput == 3 {
		evenToOneThousand()
	} else {
		fmt.Println("wrong number")
	}
	fmt.Println(message)
}
