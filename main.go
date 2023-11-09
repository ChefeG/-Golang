package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Decode(roman string) int {
	var decoder = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	if len(roman) == 0 {
		return 0
	}
	first := decoder[rune(roman[0])]
	if len(roman) == 1 {
		return first
	}
	next := decoder[rune(roman[1])]
	if next > first {
		return (next - first) + Decode(roman[2:])
	}
	return first + Decode(roman[1:])
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func main() {

	// заполнение массива
	var arr []string
	reader := bufio.NewReader(os.Stdin)
	num, _ := reader.ReadString('\n')
	num = strings.TrimSpace(num)
	arr = strings.Split(num, " ")

	// массив для проверки
	arrayNum := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// проверка на арабские числа
	var arr1 []string
	for _, elem := range arr {
		if elem == "+" || elem == "-" || elem == "*" || elem == "/" {
			arr1 = append(arr1, elem)
		}
		element, _ := strconv.Atoi(elem)
		for _, elem1 := range arrayNum {
			if element == elem1 {
				arr1 = append(arr1, elem)
			}
		}
	}

	// вычисление арабских чисел
	if len(arr1) == 3 {
		a, _ := strconv.Atoi(arr1[0])
		b, _ := strconv.Atoi(arr1[2])
		if (a > 10 || a < 1) || (b > 10 || b < 1) {
			fmt.Print("Вывод ошибки, так как одно из чисел вне диапазона (1-10)")
			arr = nil
		}
		switch {
		case arr[1] == "+":
			fmt.Print(a + b)
			arr = nil
		case arr[1] == "-":
			fmt.Print(a - b)
			arr = nil
		case arr[1] == "*":
			fmt.Print(a * b)
			arr = nil
		case arr[1] == "/":
			fmt.Print(a / b)
			arr = nil
		}
	}

	// проверка на римские числа
	var arr2 []string
	for _, elem := range arr {
		if elem == "+" || elem == "-" || elem == "*" || elem == "/" {
			arr2 = append(arr2, elem)
		}
		for _, elem1 := range arrayNum {
			if Decode(elem) == elem1 {
				arr2 = append(arr2, elem)
			}
		}
	}

	// вычисление римских чисел
	if len(arr2) == 3 {
		a := Decode(arr2[0])
		b := Decode(arr2[2])
		if (a > 10 || a < 1) || (b > 10 || b < 1) {
			fmt.Print("Вывод ошибки, так как одно из чисел вне диапазона (1-10)")
			arr = nil
		}
		switch {
		case arr[1] == "+":
			fmt.Print(integerToRoman(a + b))
		case arr[1] == "-":
			if a-b > 0 {
				fmt.Print(integerToRoman(a - b))
			} else if a-b == 0 {
				fmt.Print("Вывод ошибки, так как в римской системе нету 0.")
			} else if a-b < 0 {
				fmt.Print("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
			}
		case arr[1] == "*":
			fmt.Print(integerToRoman(a * b))
		case arr[1] == "/":
			if a/b < 1 {
				fmt.Print("Вывод ошибки, так как в римской системе нету чисел меньше 1.")
			} else {
				fmt.Print(integerToRoman(a / b))
			}
		}
	} else if len(arr1) > 3 || len(arr2) > 3 {
		fmt.Print("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).\n")
	} else if len(arr1) == 2 || len(arr2) == 2 {
		fmt.Print("Вывод ошибки, так как используются одновременно разные системы счисления.")
	} else if len(arr1) == 1 || len(arr2) == 1 {
		fmt.Print("Вывод ошибки, так как строка не является математической операцией.")
	}
}
