package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// UNIQ struct
type Uniq struct {
	data            []string
	counter         map[int]int
	result          []string
	numFields       int
	numChars        int
	caseInsensitive bool
}

// Считывает данные с STDIN или из входного файла
func (u *Uniq) readData(path string) error {
	if path == "" {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			scanner.Scan()
			if len(scanner.Text()) == 0 {
				break
			}
			u.data = append(u.data, scanner.Text())
		}
		return nil
	} else {
		f, err := os.Open(path)
		defer f.Close()
		if err != nil {
			return fmt.Errorf("Ошибка: некорректный ввод данных")
		} else {
			sc := bufio.NewScanner(f)
			for sc.Scan() {
				u.data = append(u.data, sc.Text())
			}
			return nil
		}
	}
}

// Считает, сколько раз подряд встречается очередная строка, и на основе результатов заполняет Uniq.counter
func (u *Uniq) count() {
	u.counter = make(map[int]int)
	k := 1
	s := 0
	for i := 0; i < len(u.data)-1; i++ {
		if u.caseInsensitive {
			if strings.ToLower(u.sliceData(u.data[i])) == strings.ToLower(u.sliceData(u.data[i+1])) {
				k++
			} else {
				u.counter[s] = k
				k = 1
				s = i + 1
			}
			if i == len(u.data)-2 {
				u.counter[s] = k
			}
		} else {
			if u.sliceData(u.data[i]) == u.sliceData(u.data[i+1]) {
				k++
			} else {
				u.counter[s] = k
				k = 1
				s = i + 1
			}
			if i == len(u.data)-2 {
				u.counter[s] = k
			}
		}
	}
}

// Записывает кол-во встречаний строки и саму строку в Uniq.result
func (u *Uniq) counted() {
	u.count()
	for i := 0; i < len(u.data); i++ {
		if u.counter[i] != 0 {
			u.result = append(u.result, strconv.Itoa(u.counter[i])+" "+u.data[i])
		}
	}
}

// Записывает повторяющиеся строки в Uniq.result
func (u *Uniq) duplicated() {
	u.count()
	for i := 0; i < len(u.data); i++ {
		if u.counter[i] > 1 {
			u.result = append(u.result, u.data[i])
		}
	}
}

// Записывает уникальные строки в Uniq.result
func (u *Uniq) unique() {
	u.count()
	for i := 0; i < len(u.data); i++ {
		if u.counter[i] == 1 {
			u.result = append(u.result, u.data[i])
		}
	}
}

// Записывает все строки, кроме тех, что являются дубликатами, в Uniq.result
func (u *Uniq) def() {
	u.count()
	for i := 0; i < len(u.data); i++ {
		if u.counter[i] != 0 {
			u.result = append(u.result, u.data[i])
		}
	}
}

// Делает срез строки по заданному кол-ву полей и/или символов
func (u *Uniq) sliceData(s string) string {
	str := ""
	if u.numFields > 0 {
		sArr := strings.Split(s, " ")
		for i := u.numFields; i < len(sArr); i++ {
			if i == len(sArr)-1 {
				str += sArr[i]
			} else {
				str += sArr[i] + " "
			}
		}
	}
	if str == "" {
		str = s
	}
	newStr := str
	if u.numChars > 0 {
		newStr = ""
		sArr := strings.Split(str, "")
		for i := u.numChars; i < len(sArr); i++ {
			newStr += sArr[i]
		}
	}
	return newStr
}

// Выводит данные в STDOUT
func (u *Uniq) printData(path string) error {
	if path == "" {
		fmt.Print(strings.Join(u.result, "\n"))
		return nil
	} else {
		f, err := os.Create(path)
		defer f.Close()
		if err != nil {
			return fmt.Errorf("Ошибка: некорректный файл вывода данных")
		} else {
			for i := 0; i < len(u.result); i++ { //поменять data на result
				f.WriteString(u.result[i] + "\n") //поменять data на result
			}
			return nil
		}
	}
}

// Главная функция
func main() {
	flagC := flag.Bool("c", false, "c")
	flagD := flag.Bool("d", false, "d")
	flagU := flag.Bool("u", false, "u")
	flagI := flag.Bool("i", false, "i")
	var numFields, numChars int
	flag.IntVar(&numFields, "f", 0, "f")
	flag.IntVar(&numChars, "s", 0, "s")
	flag.Parse()

	u := Uniq{numFields: numFields, numChars: numChars, caseInsensitive: *flagI}

	inputFile := ""
	outputFile := ""

	if len(os.Args) > 1 && len(os.Args[len(os.Args)-2]) >= 4 && (os.Args[len(os.Args)-2])[len(os.Args[len(os.Args)-2])-4:] == ".txt" {
		inputFile = os.Args[len(os.Args)-2]
		outputFile = os.Args[len(os.Args)-1]
	} else if len(os.Args) > 1 && len(os.Args[len(os.Args)-1]) >= 4 && (os.Args[len(os.Args)-1])[len(os.Args[len(os.Args)-1])-4:] == ".txt" {
		inputFile = os.Args[len(os.Args)-1]
	}

	u.readData(inputFile)
	if *flagC {
		u.counted()
	} else if *flagD {
		u.duplicated()
	} else if *flagU {
		u.unique()
	} else {
		u.def()
	}
	u.printData(outputFile)
}
