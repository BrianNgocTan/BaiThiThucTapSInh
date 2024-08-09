package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"unicode"
	"lab02/words"
)

func main() {
	file, err := os.Open("Example.txt")
	if err!=nil {
		fmt.Println("Error opening file: ", err)
		file.Close()
		return
	}

	f := func(r rune) bool {
		return unicode.IsSpace(r)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stringSlice := strings.FieldsFunc(line, f)
		for i:=0; i<len(stringSlice); i++ {
			stringSlice[i] = words.HandlingVulgarWord(stringSlice[i])
		}
		line = strings.Join(stringSlice, " ")
		fmt.Println(line)
	}
	file.Close()
}