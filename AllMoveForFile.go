package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("\tMINI FILE SYSTEM!")
	for {
		fmt.Println(`
		1 - Create File/Dir
		2 - Write File
		3 - Delete File 
		4 - Stop App
		`)

		fmt.Printf("Move number: ")
		scn := bufio.NewScanner(os.Stdin)
		scn.Scan()

		move, err := strconv.Atoi(scn.Text())
		if err != nil {
			log.Fatalf("Write num, no string %v", err)
		}

		nameScn := bufio.NewScanner(os.Stdin)
		switch move {
		case 1:
			fmt.Printf("Path File or Dir: ")
			nameScn.Scan()
			CreateFile(nameScn.Text())
		case 2:
			if lastFileCreatedName != "" && AccessingUser(fmt.Sprintf("Using last create file = '%v'", lastFileCreatedName)) {
				WriteFile(lastFileCreatedName)
			} else {
				fmt.Printf("$ Path File: ")
				nameScn.Scan()
				WriteFile(nameScn.Text())
			}
		case 3:
			fmt.Printf("$ Name File or dir: ")
			nameScn.Scan()
			DeleteFile(nameScn.Text())
		case 4:
			fmt.Println("$ App Stopped...")
			return
		default:
			fmt.Printf("\n\nWARNING: Write num from menu")
		}

	}
	// создание файла
	// CreateFile()
	// CreateFile2()

	// запись в файл
	// WriteFile()

	// SearchLineInFile()

	// // задание: Найти в файле .go каждую функцию и посчитать сколько строк состоит;
	// CountRowsInFuncAndFor()

	// // задание: Найти в файле .go каждую функцию и посчитать сколько строк состоит; чуть оптимизированная
	// CountFunctionLinesBalanced()

}

var lastFileCreatedName string

func AccessingUser(msg string) bool {
	scn := bufio.NewScanner(os.Stdin)
	fmt.Printf("$ %v (ignoring = n)? y/n: ", msg)
	scn.Scan()

	return strings.ToLower(scn.Text()) == "y"
}

func DeleteFile(pathFile string) {
	if err := os.Remove(pathFile); err != nil {
		fmt.Println("Fail delete fail", err)
		return
	}
	nameFail := filepath.Base(pathFile)

	fmt.Println("Access! Deleted fail", nameFail)
}

const BytesInKb = 1024

func CreateFile(name string) {
	f, err := os.Stat(name)
	flag := false
	switch {
	case err == nil:
		fmt.Println("WARNING: File is exists")
		fmt.Printf("%.2f KB\n", float64(f.Size()/BytesInKb))
		if AccessingUser("We propose to create a new dir for this file") {
			nameDir := bufio.NewScanner(os.Stdin)
			fmt.Printf("$ Name dir: ")
			nameDir.Scan()

			if err := os.Mkdir(nameDir.Text(), 0755); err != nil {
				fmt.Println("Fail created dir", err)
				return
			}
			name = nameDir.Text() + "\\" + name
			flag = true
		} else {
			fmt.Println("Operation cancelled")
			return
		}
		fallthrough // test fallthrough
	default:
		file, err := os.Create(name)
		if err != nil {
			log.Fatalf("Ошибка: %v", err)
		}
		defer file.Close()
		if flag {
			lastFileCreatedName = name
		} else {
			lastFileCreatedName = filepath.Base(name)
		}
		fmt.Println("Access! File is Created! NameFile:", name)
	}
}

// func CreateFile(name string) {
// 	if f, err := os.Stat(name); err == nil {
// 		fmt.Println("WARNING: File is exists")
// 		fmt.Printf("%.2f KB\n", float64(f.Size()/BytesInKb))
// 		if AccessingUser("We propose to create a new dir for this file") {
// 			nameDir := bufio.NewScanner(os.Stdin)
// 			fmt.Printf("$ Name dir: ")
// 			nameDir.Scan()

// 			if err := os.Mkdir(nameDir.Text(), 0755); err != nil {
// 				fmt.Println("Fail created dir", err)
// 				return
// 			}
// 			f, _ := os.Create(nameDir.Text() + "\\" + name)
// 			lastFileCreatedName = nameDir.Text() + "\\" + name
// 			defer f.Close()
// 		}
// 		return
// 	}

// 	file, err := os.Create(name)
// 	if err != nil {
// 		log.Fatalf("Ошибка: %v", err)
// 	}
// 	defer file.Close()
// 	lastFileCreatedName = filepath.Base(name)
// 	fmt.Println("Access! File is Created! NameFile: ", name)
// }

func WriteFile(nameFile string) {
	file, err := os.OpenFile(nameFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Не удалость открыть файл: %v", err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	fmt.Printf("Введите данные для записи в файл: ")
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()

	fmt.Fprintln(file, scn.Text())

	w.Flush()
}

// func CreateFile2() {
// 	err := os.WriteFile("C:/Users/rikos/OneDrive/Документы/dockerTrena/renatos.txt", []byte("reant"), 0644)
// 	if err != nil {
// 		log.Fatalf("Ошибка: %v", err)
// 	}
// }

// Поиск определенной строки в файле
// func SearchLineInFile() {
// 	file, err := os.Open("renatos.txt")
// 	if err != nil {
// 		log.Fatalf("Ошибка: %v", err)
// 	}
// 	defer file.Close()

// 	scn := bufio.NewScanner(file)
// 	linNum := 1

// 	for scn.Scan() {
// 		line := scn.Text()
// 		if strings.Contains(line, "мир") {
// 			fmt.Printf("%v: %v", linNum, line)
// 			break
// 		}
// 		linNum++
// 	}
// }

// задание: Найти в файле .go каждую функцию и посчитать сколько строк состоит;
// func CountRowsInFuncAndFor() {
// 	file, err := os.Open("../test/test.go")
// 	if err != nil {
// 		log.Fatalln("Ошибка: ", err)
// 	}
// 	defer file.Close()

// 	countRowFunc := 1
// 	// countRowFor := 1

// 	scn := bufio.NewScanner(file)

// 	stack := []rune{}

// 	flag := false
// 	for scn.Scan() {
// 		line := scn.Text()
// 		if strings.HasPrefix(line, "func ") && strings.Contains(line, "{") {
// 			flag = true
// 		}
// 		fmt.Println(line)

// 		brackets := 1
// 		if strings.Contains(line, "{") {
// 			stack = append(stack, '{')
// 			brackets++
// 		} else if strings.Contains(line, "}") {
// 			brackets--
// 			stack = stack[:len(stack)-1]
// 			if len(stack) == 0 {
// 				flag = false
// 				fmt.Println("Строк в этой функции: ", countRowFunc, brackets)
// 				countRowFunc = 0
// 			}
// 		}

// 		if flag {
// 			countRowFunc++
// 		}
// 	}

// 	fmt.Println(countRowFunc)
// }

// // задание: Найти в файле .go каждую функцию и посчитать сколько строк состоит; чуть оптимизированная версия
// func CountFunctionLinesBalanced(filename string) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatalf("Ошибка: %v", err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	braceCount := 0
// 	inFunction := false
// 	lineCount := 0
// 	currentLine := 0

// 	for scanner.Scan() {
// 		currentLine++
// 		line := scanner.Text()

// 		// Читабельно: понятная проверка начала функции
// 		if !inFunction && strings.HasPrefix(line, "func ") && strings.Contains(line, "{") {
// 			inFunction = true
// 			braceCount = 1
// 			lineCount = 1
// 			continue
// 		}

// 		// Производительно: один проход по строке для подсчета скобок
// 		if inFunction {
// 			for _, char := range line {
// 				if char == '{' {
// 					braceCount++
// 				} else if char == '}' {
// 					braceCount--
// 					if braceCount == 0 {
// 						fmt.Printf("Строка %d: функция содержит %d строк\n", currentLine, lineCount)
// 						inFunction = false
// 					}
// 				}
// 			}
// 			lineCount++
// 		}
// 	}

// 	if inFunction {
// 		fmt.Printf("Функция не закрыта: %d строк\n", lineCount)
// 	}
// }
