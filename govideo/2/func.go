package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func cal(a, b int, ops string) (int, error) {
	switch ops {
	case "+":
		return a + b, nil
	case "*":
		return a * b, nil
	default:
		return 0, fmt.Errorf("error")
	}
}

func main() {
	res, err := cal(1, 1, "+")
	if err != nil {
	} else {
		fmt.Println(res)
	}

	fmt.Println(apply(plus, 1, 31))

	fmt.Println(sum(2, 3, 4))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("this func name is %s, a=%d, b=%d\n", opName, a, b)
	return op(a, b)

}

func plus(a int, b int) int {
	return a + b
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b *int) {
	c := *a
	*a = *b
	*b = c
}
