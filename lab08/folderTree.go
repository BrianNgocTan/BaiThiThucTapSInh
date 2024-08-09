package main

import (
	"fmt"
	"os"
	"bufio"
//	"strings"
)

func folderTree(path string, degree int) {
	fields, err := os.ReadDir(path)
	if err != nil {
		fmt.Print("Error open folder:", err)
		return
	}

	for i := 0; i < degree; i++ {
		if i != degree - 1 {
      fmt.Print("   ")
		} else {
			fmt.Print("|__")
		}
	}
	var firstSlash int
	var hasSlash bool = false
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			firstSlash = i
			hasSlash = true
			break
		}
	}
	if hasSlash {
    simplePath := path[firstSlash + 1:]
	  fmt.Println(simplePath)
	} else {
		fmt.Println(path)
	}
	for i := 0; i < len(fields); i++ {	
		if fields[i].IsDir() {
			//folderTree(path + "/" + fields[i].Name(), degree + 1)
		} else {
			for j := 0; j < degree + 1; j++ {
				if j != degree {
					fmt.Print("   ")
				} else {
					fmt.Print("|__")
				}
			}
			fmt.Println(fields[i].Name())
		}
	}
	for i := 0; i < len(fields); i++ {
		if fields[i].IsDir() {
			folderTree(path + "/" + fields[i].Name(), degree + 1) // goi de quy lai ham, truong hop dung chinh la khi len(fields) == 0 hoac		                                                       
		}                                                       // fields[i] deu la chi la file
	}
}

func main() {
	fmt.Println("Nhap vao duong dan thu muc ban muon hien thi: ")
	reader := bufio.NewReader(os.Stdin)

	folderPath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error input from Stdin:", err)
		return
	}
	folderPath = folderPath[0 : len(folderPath) - 1]

	folderTree(folderPath, 0)
}