package main

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func (u User) Display() {
	fmt.Printf("ID:%d, Name:%s, Email:%s\n", u.ID, u.Name, u.Email)
}

func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func sumEven(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}

func reverse(numbers []int) []int {
	reversed := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		reversed[i] = numbers[len(numbers)-1-i]
	}
	return reversed
}

func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("division by zero")
	}
}

func countWords(text string) map[string]int {
	words := strings.Fields(text)
	wordcount := make(map[string]int)

	for _, word := range words {
		word = strings.ToLower(word)
		wordcount[word]++
	}
	return wordcount
}

func findMostFrequent(wordCount map[string]int) (string, int) {
	var mostFrequent string
	maxCount := 0
	for word, count := range wordCount {
		if count > maxCount {
			maxCount = count
			mostFrequent = word
		}
	}

	return mostFrequent, maxCount
}
func getGrade(score int) string {
	switch {
	case score >= 90 && score <= 100:
		return "A"
	case score >= 80 && score < 90:
		return "B"
	case score >= 70 && score < 80:
		return "C"
	case score >= 60 && score < 70:
		return "D"
	case score < 60:
		return "F"
	default:
		return "正しい点数を入力してください！"

	}
}

func calculateArea(length, width float64) (float64, error) {
	if length <= 0 {
		return 0, errors.New("length must be greater than 0")
	}
	if width <= 0 {
		return 0, errors.New("width must be greater than 0")
	}
	return length * width, nil
}

func main() {
	fmt.Println("Hello Go")

	result, err := calculate(10, 2, "/")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("結果：", result)

	users := []User{
		{1, "Tri", "2240788@ecc.ac.jp"},
		{2, "Tri2", "2240788@ecc.ac.jp"},
	}

	for _, user := range users {
		user.Display()
	}

	numbers := []int{1, 5, 7, 3, 2}
	fmt.Println("最大の値：", findMax(numbers))
	fmt.Println("Sum: ", sumEven(numbers))
	fmt.Println("reversed: ", reverse(numbers))

	// Day2
	text := "Go is great Go is fast Go is simple"
	wordCount := countWords(text)
	fmt.Println("文字の数", wordCount)

	word, count := findMostFrequent(wordCount)
	fmt.Printf("Từ xuất hiện nhiều nhất: '%s' xuất hiện %d lần\n", word, count)

	scores := []int{95, 85, 75, 65, 55, 105}
	for _, score := range scores {
		fmt.Printf("Score %d: Grade %s\n", score, getGrade(score))
	}

	testCases := []struct {
		length, width float64
	}{
		{5, 3},
		{-1, 3},
		{5, -2},
		{0, 10},
	}

	for _, tc := range testCases {
		area, err := calculateArea(tc.length, tc.width)
		if err != nil {
			fmt.Printf("Lỗi: %v\n", err)
		} else {
			fmt.Printf("Diện tích: %.2f\n", area)
		}
	}
}
