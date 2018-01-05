package test

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	A
}

func (a *A) Say() {
	fmt.Println("a Say()")
}

func (b *B) Say() {
	fmt.Println("b Say()")
}

func (b *B) Run() {
	fmt.Println("b Run()")
}

func TestHelloWorld() {
	//var p parser.IParser
	//p = parser.NewWandoujiaParser()
	//p.Parse(nil)
	//fmt.Println(p.GetOs())
	//fmt.Println(p.GetStoreId())
	//
	//p = parser.NewAppStoreParser()
	//p.Parse(nil)
	//fmt.Println(p.GetOs())
	//fmt.Println(p.GetStoreId())
}
