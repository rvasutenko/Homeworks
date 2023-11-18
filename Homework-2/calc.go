package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// STACK struct
type Stack struct {
	data []string
	len  int
}

// Добавляет элемент в конец Стека
func (s *Stack) Push(elem string) {
	s.data = append(s.data, elem)
	s.len++
}

// "Показывает" последний элемент Стека
func (s *Stack) Top() string {
	if s.len <= 0 {
		return ""
	} else {
		return s.data[s.len-1]
	}
}

// Забирает из Стека последний элемент
func (s *Stack) Pop() string {
	if s.len <= 0 {
		return ""
	} else {
		elem := s.data[s.len-1]
		s.data = s.data[:len(s.data)-1]
		s.len--
		return elem
	}
}

// Считывает данные с STDIN
func readData() string {
	return os.Args[len(os.Args)-1]
}

// Вычисляет конечный ответ
func calculateAnswer(str string) (float64, error) {
	operands := Stack{}
	operators := Stack{}
	for i := 0; i < len(str); i++ {
		elem := string(str[i])
		if isNumber(str[i]) {
			parseNumber(&i, str, &elem)
			operands.Push(elem)
		} else if str[i] == '(' {
			operators.Push(elem)
		} else if str[i] == ')' {
			for operators.Top() != "(" && operators.len > 0 {
				countIntermediateResult(&operands, &operators)
			}
			operators.Pop()
		} else if isOperator(str[i]) {
			if priority(elem) > priority(operators.Top()) {
				operators.Push(elem)
			} else {
				for priority(elem) <= priority(operators.Top()) && operators.len > 0 {
					countIntermediateResult(&operands, &operators)
				}
				operators.Push(elem)
			}
		} else {
			return 0, fmt.Errorf("Ошибка: введённая строка некорректна ")
		}
	}
	for operators.len > 0 {
		countIntermediateResult(&operands, &operators)
	}
	ans, _ := strconv.ParseFloat(operands.Pop(), 64)
	return ans, nil
}

// Посимвольно считывает число, пока очередной элемент не окажется арифмитическим оператором
func parseNumber(i *int, str string, elem *string) {
	j := 0
	for j = *i + 1; j < len(str); j++ {
		if isNumberContinuation(str[j]) {
			*elem += string(str[j])
		} else {
			break
		}
	}
	*i = j - 1
}

// Определяет приоритет текущей операции
func priority(operator string) int {
	switch operator {
	case "^":
		return 3
	case "*":
		return 2
	case "/":
		return 2
	case "+":
		return 1
	case "-":
		return 1
	default:
		return -1
	}
}

// Вычисляет промежуточный результат
func countIntermediateResult(operands, operators *Stack) {
	b, _ := strconv.ParseFloat(operands.Pop(), 64)
	a, _ := strconv.ParseFloat(operands.Pop(), 64)
	operator := operators.Pop()
	var res float64
	switch operator {
	case "^":
		res = math.Pow(a, b)
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "+":
		res = a + b
	case "-":
		res = a - b
	}
	operands.Push(strconv.FormatFloat(res, 'E', -1, 64))
}

// Определяет, является ли текущий символ числом
func isNumber(index uint8) bool {
	return index >= '0' && index <= '9'
}

// Определяет, является ли текущий символ числом / составным элементом числа
func isNumberContinuation(index uint8) bool {
	return (index >= '0' && index <= '9') || index == '.'
}

// Определяет, является ли текущий символ оператором
func isOperator(index uint8) bool {
	return index == '^' || index == '*' || index == '/' || index == '+' || index == '-'
}

// Главная функция
func main() {
	inputStr := readData()
	ans, err := calculateAnswer(inputStr)
	if err == nil {
		fmt.Println(ans)
	} else {
		fmt.Println(err)
	}
}
