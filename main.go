package main

import (
	"fmt"
	"strings"
)

// struct
type person struct {
	name    string
	age     int
	favFood []string
}

// 조건문
func canIDrink(age int) bool {
	// variable expression
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

// switch - if 처럼 varialbe expression 사용가능 함
func eatKimch(age int) bool {
	switch {
	case age < 18:
		return false
	case age > 18:
		return true
	}
	return false
}

// loop는 for밖에 없음
func superAdd(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	total := 0
	// range를 사용하면 무조건 인덱스를 줌
	for _, number := range numbers {
		total += number
	}
	return total
}

// ... 사용하면 무제한으로 argument 를 받을 수 있다.
func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked return, defer ( 함수가 끝난 후 실행되는 코드)
func lenAndUpperNaked(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// go 함수는 여러가지를 반환 할 수 있다.
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	// 상수
	const name string = "simon"
	// 변수
	var name2 string = "simon2"
	// := 자동으로 타입을 추측해서 입력해준다...
	name3 := "simon3"
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("simon")
	// ingnored value ( _ 사용해서 무시 가능)
	totalLength2, _ := lenAndUpper("simon")
	fmt.Println(totalLength, upperName)
	fmt.Println(totalLength2)

	repeatMe("simon", "lynn", "kimch", "nothing")

	// naked function, defer
	t, u := lenAndUpperNaked("dorothy")
	fmt.Println(t, u)

	// loop
	total := superAdd(1, 2, 3, 4, 56, 7)
	fmt.Println(total)

	// 조건문 👇
	fmt.Println(canIDrink(16))

	// pointer
	a := 2
	b := &a
	a = 5
	fmt.Println(&a, b)
	// 이 시점에서 b는 a의 메모리 주소를 바라보고 있고
	// * 을 사용해서 그 주소의 값을 확인할 수 있다.
	fmt.Println(*b)
	// 반대로 b가 가지고 있는 메모리 주소로 a 의 값을 변경할 수 있다.
	*b = 20
	fmt.Println(a)

	// array
	// 배열의 길이를 지정해 줘야 함
	names := [5]string{"simon", "me", "you", "I"}
	names[4] = "lalal"
	fmt.Println(names)
	// slice는 무한하게 늘어나는 배열이다
	sliceName := []string{"slice", "array"}
	fmt.Println(sliceName)
	// slice에 추가하는 함수 - 항상 반환한다
	sliceName = append(sliceName, "fly")
	fmt.Println(sliceName)

	// map
	// map[key]value{ ~ }
	simon := map[string]string{"name": "simon", "age": "12"}
	fmt.Println(simon)
	for key, value := range simon {
		fmt.Println(key, value)
	}

	// struct
	favFood := []string{"ham", "bab"}
	// 2가지 다 가능
	// people := person{"simon", 20, favFood}
	people := person{name: "simon", age: 20, favFood: favFood}
	fmt.Println(people)
}
