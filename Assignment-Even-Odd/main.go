package main

import "fmt"


func main()  {
	
	numbers := []int{0,1,2,3,4,5,6,7,8,9,10}

	fmt.Println(num)

	for _, number := range numbers{
		if number % 2 == 0 {
			fmt.Println(number, "is EVEN")
		} else {
			fmt.Println(number, "is odd")
		}
	}
	

}