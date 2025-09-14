package main

import "fmt"

func main() {
	fmt.Println("Hello this is the first file in this learning go directory.")
	fmt.Println("We are going to learn more and more about go as we progress")

	packages1()
	packages2()
	packages3()

	sum := add(3, 7)
	fmt.Println("Sum is", sum)

	sum2 := addbutbetter(3, 11)
	fmt.Println("Sum is", sum2)

	a, b := swap("learn", "fast")
	fmt.Println(a, b)

	fmt.Print(split(17))

}
