package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func readFileAtLineY(nameFile string, Y int) error {
	file, err := os.Open(nameFile)
	if err != nil {
		file.Close()
		return err
	}

	scanner := bufio.NewScanner(file)
	var count int = 0
	var hasYLine bool = false
	for scanner.Scan() {
		count++
		if count == Y {
			hasYLine = true
			fmt.Println("Noi dung dong thu", Y, ":", scanner.Text())
			break
		}		
	}
	if hasYLine {

	} else {
		err = scanner.Err() //viec doc cac dong trong file co the xay ra su co khi Y qua lon va file qua nhieu dong
		if err != nil {
			return err
		}

		fmt.Println("File ", nameFile, " don't have line with order ", Y) //viec doc file khong co issue, chac chan file khong co dong thu Y
	}

	file.Close()

	return err
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var name, y string
	var err error
	fmt.Println("Nhap vao ten file muon mo: ")
	name, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error input from Stdin: ", err)
		return
	}
	name = name[: len(name) -1] //loai bo ki tu '\n'
	fmt.Println("Nhap thu tu cua dong muon doc: ")
	y, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Err input from Stdin: ", err)
		return
	}
	y = y[: len(y) -1] // loai bo ki tu '\n'
	var yNumber int
	yNumber, err = strconv.Atoi(y)
	if err != nil {
		fmt.Println("Error convert string to number: ", err)
		return
	}
	
	err = readFileAtLineY(name, yNumber)
	if err != nil {
		fmt.Println("Error when open and read file: ", err)
	}
}