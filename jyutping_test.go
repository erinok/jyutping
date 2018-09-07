package jyutping

import "fmt"

func ExampleConvert() {
	fmt.Println(Convert("辦唔到。"))
	fmt.Println(Convert("我證件唔見晒仲辦唔辦到？"))
	fmt.Println(Convert("hello我 funny證件唔見晒😇仲?!?! funky times...辦唔辦到？"))
	// Output:
	// baan6 m4 dou3。
	// ngo5 zing3 gin2 m4 gin3 saai3 zung6 baan6 m4 baan6 dou3？
	// hello ngo5 funny zing3 gin2 m4 gin3 saai3😇 zung6?!?! funky times... baan6 m4 baan6 dou3？
}
