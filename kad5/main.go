package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// JKad05D1
	fmt.Println("-----JKad05D1-----")
	// scanner := bufio.NewScanner(os.Stdin)
	var n1 int
	fmt.Print("ひとつめの整数を入力してください＞")
	fmt.Scan(&n1)

	var n2 int
	fmt.Print("ふたつめの整数を入力してください＞")
	fmt.Scan(&n2)

	if n1 > n2 {
		fmt.Println("ふたつの数字で大きい方は ", n1, "です。")
	} else {
		fmt.Println("ふたつの整数で大きい方は", n2, "です。")
	}

	// JKad05D2
	fmt.Println("\n-----JKad05D2-----")
	fmt.Println("こんにちは！\nわたしはスーパーティーチャー、ECC エクセレントよ！ ")
	class := [3]string{"4 年制コース", "3 年制コース", "2 年制コース "}
	fmt.Print("あなたのクラスを教えてね！（0：IE、1：SK、2：SE）＞")
	var classNumber int
	fmt.Scan(&classNumber)
	if classNumber >= 0 && classNumber <= 2 {
		fmt.Println(class[classNumber], "コースなのね！")
	} else {
		fmt.Println("無効な番号です。0、1、2 を入力してね！")
	}

	// JKad05X
	fmt.Println("\n-----JKad05X-----")
	fmt.Println("こんにちは！")
	fmt.Println("わたしはスーパーティーチャー、ECC エクセレントよ！")
	fmt.Print("あなたのクラスを教えてね！（0：IE、1：SK、2：SE）＞")
	var classNumberX int
	fmt.Scan(&classNumberX)
	if classNumberX == 0 {
		fmt.Println(class[classNumberX], "コースなのね！")
	} else if classNumberX == 1 {
		fmt.Println(class[classNumberX], "コースなのね！")
	} else if classNumberX == 2 {
		fmt.Println(class[classNumberX], "コースなのね！")
	} else {
		fmt.Println("わからないのね！")
	}

	// JKad05C1
	fmt.Println("\n-----JKad05C1-----")
	fmt.Println("ひとり一つずつリンゴを食べます！")
	var apples int
	fmt.Print("リンゴの数を入力してください＞")
	fmt.Scan(&apples)

	var people int
	fmt.Print("食べる人の数を入力してください＞")
	fmt.Scan(&people)

	if apples >= people {
		fmt.Println("ひとり一つずつリンゴを食べた！")
		fmt.Println("残ったリンゴは", apples-people, "個です！")
	} else {
		fmt.Println("リンゴの数が足りません！")
	}

	// JKad05B
	fmt.Println("\n-----JKad05B-----")
	fmt.Println("２つのサイコロを振ります！")
	fmt.Print("出た目の和が偶数か奇数か予想してください（0：偶数、1：奇数）＞")
	var guess int
	fmt.Scan(&guess)

	dice1 := rand.Intn(6) + 1
	dice2 := rand.Intn(6) + 1
	sum := dice1 + dice2
	fmt.Println("ひとつめは", dice1, "！")
	fmt.Println("ふたつめは", dice2, "！")
	fmt.Println("合わせて", sum, "！")

	var isEven int
	if sum%2 == 0 {
		fmt.Println("偶数でした！")
		isEven = 0
	} else {
		fmt.Println("奇数でした！")
		isEven = 1
	}

	if guess == isEven {
		fmt.Println("あなたの勝ち！")
	} else {
		fmt.Println("あなたの負け！")
	}

	// JKad05C2
	fmt.Println("\n-----JKad05C2-----")
	fmt.Println("乱数を発生させます！")
	n := rand.Intn(10)
	fmt.Println(n, "が出ました！")
	if n%2 == 0 {
		fmt.Println("偶数でした！")
	} else {
		fmt.Println("奇数でした！")
	}

	// JKad05A
	fmt.Println("\n-----JKad05A-----")
	names := [5]string{"のび太", "しずかちゃん", "ジャイアン", "スネ夫", "出木杉くん"}
	var scores [5]int
	for i := 0; i < 5; i++ {
		fmt.Print(names[i], "の点数を入力してください＞")
		fmt.Scan(&scores[i])
	}

	max := scores[0]
	for i := 1; i < 5; i++ {
		if scores[i] > max {
			max = scores[i]
		}
	}
	fmt.Println("一番高い点数は", max, "点です！")

	// JKad05S
	fmt.Println("\n-----JKad05S-----")
	fmt.Println("こんにちは！")
	fmt.Println("わたしは占いマシーンの ECC1000 よ！")
	fmt.Println("あなたのことを占ってあげるわ！よろしくね")
	fmt.Println()

	var name string
	fmt.Print("名前は何ていうの？＞")
	fmt.Scan(&name)
	var age int
	fmt.Print("年齢はいくつ？＞")
	fmt.Scan(&age)
	fmt.Println()

	fmt.Println(name, "さん、こんにちは！")
	fmt.Println("あなたは", age, "歳なんですね！")
	fmt.Println()

	loveScore := rand.Intn(101)
	moneyScore := rand.Intn(101)
	overallScore := rand.Intn(101)

	fmt.Println(name, "さんの今日の運勢は")
	fmt.Println("ラブ運 ", loveScore)
	fmt.Println("金銭運 ", moneyScore)
	fmt.Println("全体運 ", overallScore)
	fmt.Println()

	bestName := "ラブ運"
	bestScore := loveScore
	if moneyScore > bestScore {
		bestName = "金銭運"
		bestScore = moneyScore
	}
	if overallScore > bestScore {
		bestName = "全体運"
		bestScore = overallScore
	}
	fmt.Println("一番高いのは" + bestName + "の" + fmt.Sprintf("%d", bestScore) + " ね！")

}
