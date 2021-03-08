package main

import "fmt"

func main() {
	const name string = "simon"
	var name2 string = "simon2"
	// := 자동으로 타입을 추측해서 입력해준다.
	name3 := "simon3"
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
}
