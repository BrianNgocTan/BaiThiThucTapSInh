package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"unicode"
)

func checkKeyInLine(line string, key string) bool {
	var f = func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := strings.FieldsFunc(line, f)
	//copyWords la ban sao cua words, khac o cho cac string trong copyWords la cac string trong words duoc chuyen cac chu in hoa thanh in thuong
	var copyWords []string
	for i := range words {
		var word []byte
		for j := range words[i] {
			if 'A' <= words[i][j] && words[i][j] <= 'Z' {
				word = append(word, words[i][j] + 32)
			} else {
				word = append(word, words[i][j])
			}
		}
		copyWords = append(copyWords, string(word))
	}
	//copyKey la ban sao cu key, chi khac o cho cac chu in hoa trong key duoc chuyen ve chu thuong trong copykey
	var copyKey string
	var s[]byte
	for i := range key {		
		if 'A' <= key[i] && key[i] <= 'Z' {
			s = append(s, key[i] + 32)
		} else {
			s = append(s, key[i])
		}
	}
	copyKey = string(s)

	var hasKey = false
	for _, word := range copyWords {
		if word == copyKey {
			hasKey = true
			break
		}
	}

	return hasKey
}

func main() {
	var pathName, key string // := "/home/tan/Documents" home/tan/HocTap/Golang
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Nhap duong dan thu muc tu ban phim: ")
	pathName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error input from Stdin:", err)
		return
	}
	pathName = pathName[: len(pathName) - 1] //loai bo di ki tu '\n'
	fmt.Println("Nhap tu khoa muon tim kiem: ")
	key, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error input from Stdin:", err)
		return
	}
	key = key[: len(key) -1] // loai bo di ki tu '\n'

	files, err1 := os.ReadDir(pathName)
	if err1 != nil {
		fmt.Println("Error reading directory", err1)
		return
	}
	var txtFiles []string //slice luu cac ten co duoi .txt cua cac file trong thu muc
	for _, file := range files {
		if file.IsDir() {
			//fmt.Println(file.Name())
		} else {
			if strings.HasSuffix(file.Name(), ".txt") {
				txtFiles = append(txtFiles, pathName + "/" + file.Name())
			}
		}	
	}

	if len(txtFiles) == 0 {
		fmt.Println("Khong co file .txt trong thu muc da nhap")
		return
	}
	var hasKeyFile bool = false
	for i := range txtFiles {
		file, err2 := os.Open(txtFiles[i])
		if err2 != nil {
			fmt.Println("Error open file", txtFiles[i], err2)
			file.Close()
			return
		}

		scanner := bufio.NewScanner(file)
		var hasKeyLines []string
		var count int = 0
		for scanner.Scan() {
			count++
			line := scanner.Text()
			if checkKeyInLine(line, key) {
				line = ": " + line
				hasKeyLines = append(hasKeyLines, fmt.Sprintf("%d%s", count, line))
				hasKeyFile = true
			}
		}
		if len(hasKeyLines) > 0 {
			var relativeFileName string = txtFiles[i][len(pathName) + 1:]
			fmt.Println(relativeFileName + ": ")
			for _, line := range hasKeyLines {
				fmt.Println(line)
			}
			fmt.Println()
		}	
		err3 := scanner.Err()
		if err3 != nil {
			fmt.Println("Error read file ", txtFiles[i], " by scanner:", err3)
		}
		file.Close()
	}
	if hasKeyFile == false {
		fmt.Println("khong ton tai file .txt co chua key da nhap")
	}
}