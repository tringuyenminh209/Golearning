package main

import (
	"errors"
	"fmt"
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
}
