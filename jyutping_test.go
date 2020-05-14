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

func ExampleConvert7() {
	p("《權力遊戲》、《紙牌屋》")
	// Output:
	// "kyun4 lik6 jau4 hei3", "zi2 paai2 uk1"
}

func ExampleConvert8() {
	p("“Owl Likes”")
	// Output:
	// "Owl Likes"
}

func ExampleConvert9() {
	p("“啤酒為什麼會不好喝？因為雪藏得不夠冷”")
	// Output:
	// "be1 zau2 wai6 sam6 mo1 wui2 bat1 hou2 hot3? jan1 wai6 syut3 cong4 dak1 bat1 gau3 laang5"
	
}

func cc(s string) {
	fmt.Println(ColorizeChars(s))
}

func ExampleColorizeChars1() {
	cc("啤酒")
	// Output:
	// <span class="tone1">啤</span><span class="tone2">酒</span>
}

func ExampleColorizeChars2() {
	cc("啤酒為hi得")
	// Output:
	// <span class="tone1">啤</span><span class="tone2">酒</span><span class="tone6">為</span>hi<span class="tone1">得</span>
}

func ExampleColorizeChars3() {
	cc("《權力遊戲》、《紙牌屋》")
	// Output:
	// 《<span class="tone4">權</span><span class="tone6">力</span><span class="tone4">遊</span><span class="tone3">戲</span>》、《<span class="tone2">紙</span><span class="tone2">牌</span><span class="tone1">屋</span>》
}

func ExampleColorizeJP() {
	fmt.Println(colorizeJP("jat1 go3 jyut6"))
	// Output:
	// <span class="tone1">jat</span> <span class="tone3">go</span> <span class="tone6">jyut</span>
}
