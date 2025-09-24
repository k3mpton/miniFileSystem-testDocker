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
			fmt.Printf("\n\n! WARNING: Write num from menu")
		}

	}
}

var lastFileCreatedName string

func AccessingUser(msg string) bool {
	scn := bufio.NewScanner(os.Stdin)
	fmt.Printf("$ %v ( ignoring = n )? y/n: ", msg)
	scn.Scan()

	return strings.ToLower(strings.TrimSpace(scn.Text())) == "y"
}

func DeleteFile(pathFile string) {
	if err := os.Remove(pathFile); err != nil {
		fmt.Println("! Fail delete fail", err)
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
		fallthrough
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

func WriteFile(nameFile string) {
	file, err := os.OpenFile(nameFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("! Fail open fail: %v", err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	fmt.Printf("$ Write data in fail: ")
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()

	fmt.Fprintln(file, scn.Text())

	w.Flush()
}
