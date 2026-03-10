package main

import "fmt"

func main() {
	// JKad06D
	fmt.Println("-----JKad06D-----")
	var apple int
	fmt.Print("リンゴの数を入力してください＞")
	fmt.Scan(&apple)

	var person int
	fmt.Print("食べる人の数を入力してください＞")
	fmt.Scan(&person)

	for apple >= person {
		apple -= person
		fmt.Printf("ひとり一つずつリンゴを食べた！残ったリンゴは%d個です！\n", apple)
	}

	// JKad06C1
	fmt.Println("\n-----JKad06C1-----")
	var n int
	fmt.Print("1 より大きい整数を入力してください＞")
	fmt.Scan(&n)
	for ; n >= 1; n-- {
		fmt.Println(n)
	}

	//JKad06C2
	fmt.Println("\n-----JKad06C2-----")
	var n2 int
	fmt.Print("1 より大きい整数を入力してください＞")
	fmt.Scan(&n2)
	for i := 1; i <= n2; i++ {
		fmt.Println(i)
	}

	//JKad06B
	fmt.Println("\n-----JKad06B-----")
	var n3 int
	fmt.Print("1 より大きい整数を入力してください＞")
	fmt.Scan(&n3)
	sum := 0
	fmt.Printf("1 から%d まで加算します！", n3)
	for i := 1; i <= n3; i++ {
		sum += i
	}
	fmt.Printf("\n合計は%dです！！", sum)

	//JKad06A1
	fmt.Println("\n-----JKad06A1-----")
	var n4 int
	fmt.Print("整数1 を入力してください＞")
	fmt.Scan(&n4)
	var n5 int
	fmt.Print("整数2 を入力してください＞")
	fmt.Scan(&n5)
	n6 := n4 / n5
	n7 := n4 % n5
	fmt.Printf("%d÷%d を計算します！", n6, n7)
	fmt.Printf("商は%d、余りは%d です！", n6, n7)

}
