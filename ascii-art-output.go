package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		file, _ := os.Open("standard.txt")
		fileVal := ScanFile(file)
		arg := os.Args[1]
		for _, v := range os.Args[1:1] {
			arg += " " + v
		}
		narg := strings.Split(arg, "\\n")
		for _, v := range narg {
			printLetter(string(v), fileVal)
		}
		return
	} else if os.Args[2] == "thinkertoy" {
		file, _ := os.Open("thinkertoy.txt")
		fileVal := ScanFile(file)
		arg := os.Args[1]
		for _, v := range os.Args[2:2] {
			arg += " " + v
		}
		narg := strings.Split(arg, "\\n")
		for _, v := range narg {
			printLetter(string(v), fileVal)
		}
	} else if os.Args[2] == "shadow" {
		file, _ := os.Open("shadow.txt")
		fileVal := ScanFile(file)
		arg := os.Args[1]
		for _, v := range os.Args[2:2] {
			arg += " " + v
		}
		narg := strings.Split(arg, "\\n")
		for _, v := range narg {
			printLetter(string(v), fileVal)
		}
	} else if os.Args[2] == "standard" {
		file, _ := os.Open("standard.txt")
		fileVal := ScanFile(file)
		arg := os.Args[1]
		for _, v := range os.Args[2:2] {
			arg += " " + v
		}
		narg := strings.Split(arg, "\\n")
		for _, v := range narg {
			printLetter(string(v), fileVal)
		}
	} else {
		fmt.Println("ЭТО НЕ БАГ, А ФИЧА")
	}

}

func printLetter(s string, fileVal []string) {
	output := "--output="
	fileName := ""
	var help string
	arg3 := os.Args[3]

	for i := 1; i <= 8; i++ {
		for _, arg := range s {
			indexBase := int(rune(arg)-32) * 9
			help += string(fileVal[indexBase+i])
		}
		help += "\n"
	}

	if output == arg3[:9] {
		fileName = arg3[9:]
	} else {
		fmt.Println("Братан, проверь свой аутпут")
		return
	}

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(help)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "элементов записано в файл. Проверь по cat, если не веришь мне на слово!")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func ScanFile(arg *os.File) []string {
	var fileValue []string
	scanner := bufio.NewScanner(arg)
	for scanner.Scan() {
		lines := scanner.Text()
		fileValue = append(fileValue, lines)
	}
	return fileValue
}
