package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("Example.txt")
	if err!=nil {
		fmt.Println("Error opening file: ", err)
		file.Close()
		return
	}

	content, err1 := io.ReadAll(file)
	if err1!=nil {
		fmt.Println("Error reading file: ", err1)
		file.Close()
		return
	} 

  fmt.Println(string(content))
}