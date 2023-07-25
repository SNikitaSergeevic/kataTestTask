package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println(calculate(input))
}

const (
	NotExistNumber     = "Вывод ошибки, недопустимое число."
	FailedNotation     = "Вывод ошибки, так как используются одновременно разные системы счисления."
	SignAndCountFailed = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	RomeNotNegative    = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	RomeMoreOne        = "Вывод ошибки, результат выражения с римскими числами не может быть меньше единицы."
	NotMath            = "Вывод ошибки, так как строка не является математической операцией."
	FailedOneAndTen    = "Вывод ошибки, введенные числа должны быть от 1 до 10"
)

func calculate(expr string) string {

	expr2 := strings.Fields(expr)
	res := ";)"

	if len(expr2) > 3 {
		return SignAndCountFailed
	} else if len(expr2) < 3 {
		return NotMath
	}

	valDict := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	checkRome := false

	val1, err1 := strconv.Atoi(expr2[0])
	if err1 != nil {
		_, ok := valDict[expr2[0]]
		if ok {
			val1 = valDict[expr2[0]]
			checkRome = true
		} else {
			res = NotExistNumber
			return res
		}
	}

	val2, err2 := strconv.Atoi(expr2[2])
	if err2 != nil && checkRome {
		_, ok := valDict[expr2[2]]
		if ok {
			val2 = valDict[expr2[2]]
		} else {
			res = NotExistNumber
			return res
		}
	} else if checkRome || err2 != nil {
		return FailedNotation
	}

	if val1 < 1 || val1 > 10 || val2 < 1 || val2 > 10 {
		return FailedOneAndTen
	}

	sign := expr2[1]

	switch sign {
	case "+":
		if checkRome {
			res = romeTr(val1 + val2)
		}
		res = fmt.Sprint(val1 + val2)
	case "-":
		if checkRome && (val1-val2) < 1 {
			res = RomeMoreOne
		}
		if checkRome {
			res = romeTr(val1 - val2)
		}
		res = fmt.Sprint(val1 - val2)
	case "/":
		if checkRome && (val1/val2) < 1 {
			res = RomeMoreOne
		}
		if checkRome {
			res = romeTr(val1 / val2)
		}
		res = fmt.Sprint(val1 / val2)
	case "*":
		if checkRome {
			res = romeTr(val1 * val2)
		}
		res = fmt.Sprint(val1 * val2)
	default:
		res = SignAndCountFailed
	}

	return res
}

func romeTr(val int) string {
	valDict := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X"}
	res := ""

	tmp := val / 100
	res += strings.Repeat("C", tmp)
	tmp = val % 100

	tmp1 := tmp / 50
	res += strings.Repeat("L", tmp1)
	tmp1 = tmp % 50

	tmp = tmp1 / 10
	res += strings.Repeat("X", tmp)

	tmp = tmp1 % 10
	res += valDict[tmp]

	return res
}
