package test

import "fmt"

type TestClassA struct {

}

type TestClassB struct {
	TestClassA
}

type IShow interface {
	show()
}

func TestClass() {
	//db.Open()
	//var b bean.AppBean
	//b.AppId = "com.xx.cc"
	//b.StoreId = "wandoujia"
	//db.ReplaceApp(&b)
	//
	//var c bean.CategoryBean
	//c.Name = "c1"
	//c.SuperName = "c2"
	//db.ReplaceCategory(&c)
	//db.Close()
	var b = new(TestClassB)
	b.toShow()
	b.show()
}

func (p *TestClassA) toShow() {
	fmt.Println("toShow()")
	p.show()
}

func (p *TestClassA) show() {
	fmt.Println("TestClassA show()")
}

//func (p *TestClassB) show() {
//	fmt.Println("TestClassB show()")
//}