package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// JKad04D
	fmt.Println("-----JKad04D------")
	message1 := "こんにちは！"
	fmt.Println(message1)
	message2 := "今日もよい天気ですね！"
	message3 := message1 + message2
	fmt.Println(message3)

	// JKad04C
	fmt.Println("\n-----JKad04C-----")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("あなたのお名前は？＞")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	fmt.Println(name + "さん、こんにちは！")

	fmt.Println("年齢はいくつです？＞")
	var age int
	fmt.Scan(&age)
	fmt.Println(age, "歳なのね！")

	// JKad04B
	fmt.Println("\n-----JKad04B-----")
	fmt.Print("いくつまで足し合わせますか？")
	var n int
	fmt.Scan(&n)
	sum := n * (n + 1) / 2
	fmt.Printf("1から%dまで足し合わせると・・・%dになりました！\n", n, sum)

	// JKad04A1
	fmt.Println("\n-----JKad04A1-----")
	fmt.Print("教室番号を入力してください＞")
	var numclass int
	fmt.Scan(&numclass)
	fmt.Println("一の位\t", numclass%10)
	fmt.Println("十の位\t", (numclass/10)%10)
	fmt.Println("百の位\t", (numclass/100)%10)
	fmt.Println("千の位\t", (numclass/1000)%10)

	// JKad04A2
	fmt.Println("\n-----JKad04A2-----")
	normal := ""
	reverse := ""
	var str string

	for i := 1; i <= 4; i++ {
		fmt.Printf("文字列%d＞", i)
		fmt.Scan(&str)
		normal = normal + str
		reverse = str + reverse
	}

	fmt.Println("入力順：" + normal)
	fmt.Println("逆　順：" + reverse)

	// JKad04S1
	fmt.Println("\n-----JKad04S1-----")
	fmt.Print("0から255までの整数を入力してください＞")
	var n1 int
	fmt.Scan(&n1)
	fmt.Println("最下位ビットから順に表示します！")
	for i := 0; i < 8; i++ {
		fmt.Println(n & 1)
		n >>= 1
	}

	// JKad04S2
	fmt.Println("\n-----JKad04S2-----")
	fmt.Print("0から255までの整数を入力してください＞")
	var n2 int
	fmt.Scan(&n2)
	fmt.Println("最上位ビットから順に表示します！")
	result := ""
	for i := 0; i < 8; i++ {
		result = fmt.Sprintf("%d", n&2) + result // 先頭に追加
		n2 >>= 1
	}
	fmt.Println(result)

	// JKad04X
	fmt.Println("\n-----JKad04X-----")
	fmt.Print("0から65535までの整数を入力してください＞")
	var n3 int
	fmt.Scan(&n3)
	fmt.Println("最上位ビットから順に表示します！")

	power := 32768 // 2^15
	// 先頭の0をスキップ（Javaのwhile文に相当）
	for power > n3 && power > 1 {
		power /= 2
	}
	// ビットを表示
	result1 := ""
	for power >= 1 {
		if n3&power != 0 {
			result1 += "1"
		} else {
			result1 += "0"
		}
		power /= 2
	}
	fmt.Println(result1)
}
