package test

import (
	"flag"
	"fmt"
)

func TestCmd() {
	//for _, value := range os.Args[1:] {
	//	fmt.Println(value)
	//}
	var u string
	flag.StringVar(&u, "u", "default u", "")
	flag.Parse()
	//flags := flag.Args()
	//for _, v := range flags {
	//	fmt.Println("v", v)
	//}
	fmt.Println("u", u)
}
