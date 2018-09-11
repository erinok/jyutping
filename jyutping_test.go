package jyutping

import "fmt"

func p(s string) {
	fmt.Println(Convert(s))
}

func ExampleConvert() {
	p("辦唔到。")
	p("我證件唔見晒仲辦唔辦到？")
	p("hello我funny證件唔見晒😇仲?!?! funky times are here...辦唔辦到？")
	// Output:
	// baan6 m4 dou3.
	// ngo5 zing3 gin2 m4 gin3 saai3 zung6 baan6 m4 baan6 dou3?
	// hello ngo5 funny zing3 gin2 m4 gin3 saai3😇 zung6?!?! funky times are here... baan6 m4 baan6 dou3?
}

func ExampleConvert2() {
	p("腳，來")
	// Output:
	// goek3, loi4
}

func ExampleConvert3() {
	p("他們一開口會說：“阿芝妳的眼角開始有點下墜”")
	// Output:
	// taa1 mun4 jat1 hoi1 hau2 wui2 seoi3: "aa3 zi1 nei5 dik1 ngaan5 gok3 hoi1 ci2 jau5 dim2 haa6 zeoi6"
}

func ExampleConvert4() {
	p("例如《全部都係雞》")
	// Output:
	// lai6 jyu4 "cyun4 bou6 dou1 hai6 gai1"
}

func ExampleConvert5() {
	p("字幕翻譯：李淽芊")
	// Output:
	// zi6 mok6 faan1 jik6: lei5 zi2 cin1
}


func ExampleConvert6() {
	p("《梅花》的背景是1970年代的")
	// Output:
	// "mui4 faa1" dik1 bui3 ging2 si6 1970 nin4 doi6 dik1
}

	
