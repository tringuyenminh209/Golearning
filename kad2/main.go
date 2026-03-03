package main

import "fmt"

func main() {

	//JKad02A
	apple := 20
	person := 3
	fmt.Println("リンゴが", apple, "個ありました！\n", apple)
	fmt.Println("人間が", person, "人やってきました！\n", person)
	fmt.Println("それぞれが食べたリンゴの数は\n", apple/person)
	fmt.Println("残ったリンゴの数は\n", apple%person)

	//JKad02S
	x := 2
	y := 7
	fmt.Println("xの値は\n", x)
	fmt.Println("yの値は\n", y)
	fmt.Println("*** x とyの値を入れ替えます！ ***")
	tmp := x
	x = y
	y = tmp
	fmt.Println("xの値は\n", x)
	fmt.Println("yの値は\n", y)

	//JKad02X1
	n := 10
	sum := 0
	fmt.Println("nの値\n", n)
	fmt.Println("*** 1 からnまで加算します！ *** ")
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("結果は\n", sum)
}
