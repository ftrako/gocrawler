package test

import "fmt"

type TestEnumType int

const (
	TestEnum_A = iota
	TestEnum_B
	TestEnum_C = 10
	TestEnum_D = 10 + iota
	TestEnum_E
)

func TestEnum() {
	fmt.Println("A", TestEnum_A)
	fmt.Println("B", TestEnum_B)
	fmt.Println("C", TestEnum_C)
	fmt.Println("D", TestEnum_D)
	fmt.Println("E", TestEnum_E)
}
