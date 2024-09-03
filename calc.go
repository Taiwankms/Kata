package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	romeToArab := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"XX": 20, "XXX": 30, "XL": 40, "L": 50,
		"LX": 60, "LXX": 70, "LXXX": 80, "XC": 90, "C": 100,
	}
	arabToRome := func(number int) string {
		if number <= 0 {
			return ""
		}

		var result strings.Builder

		for number >= 100 {
			result.WriteString("C")
			number -= 100
		}
		if number >= 90 {
			result.WriteString("XC")
			number -= 90
		}
		if number >= 50 {
			result.WriteString("L")
			number -= 50
		}
		if number >= 40 {
			result.WriteString("XL")
			number -= 40
		}
		for number >= 10 {
			result.WriteString("X")
			number -= 10
		}
		if number >= 9 {
			result.WriteString("IX")
			number -= 9
		}
		if number >= 5 {
			result.WriteString("V")
			number -= 5
		}
		if number >= 4 {
			result.WriteString("IV")
			number -= 4
		}
		for number >= 1 {
			result.WriteString("I")
			number--
		}
		return result.String()
	}
	isValidOperation := func(op string) bool {
		switch op {
		case "+", "-", "*", "/":
			return true
		default:
			return false
		}
	}
	var result int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение вида 2 + 2, или если ты римлянин то II + II(числа должны быть от 1 до 10 или от I до X): ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	
	parts := strings.Split(text, " ")
	if len(parts) != 3 {
		fmt.Println("Паника, формат ввода совершенно не тот, надо два числа и оператор между ними, и все через пробел!")
		return
	}
	aStr, bStr := parts[0], parts[2]
	operation := parts[1]
	if !isValidOperation(operation) {
		fmt.Println("Паника, это что угодно, но точно не математическая операция!")
		return
	}
	var a, b int
	var err1, err2 error
	var isRome bool
	a, err1 = strconv.Atoi(aStr)
	b, err2 = strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		isRome = true
		a = romeToArab[aStr]
		b = romeToArab[bStr]
		if a == 0 || b == 0 || a > 10 || b > 10 {
			fmt.Println("Паника, введено что то не то!")
			return
		}
	} else {
		if (a < 1 || a > 10) || (b < 1 || b > 10) {
			fmt.Println("Паника, как минимум 1 число не входит в диапазон от 1 до 10!")
			return
		}
	}

	switch operation {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			result = a / b			
		default:
			fmt.Println("Что то пошло не так, попробуйте еще раз")
			return
		}
		if isRome {
			if result <= 0 {
				fmt.Println("Паника: настоящие римляне не знают что такое ноль и отрицательное число!")
				return
			}
			fmt.Println("Ответ: ", arabToRome(result))
		} else {
		fmt.Println("Ответ:", result)
		}
}

