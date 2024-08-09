package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("Example.txt")
	if err!=nil {
		fmt.Println("Error creating file: ", err)
		file.Close()
		return
	}

	content := "Hoc lap trinh golang"
	_, err = file.WriteString(content)
	if err!=nil {
		fmt.Println("Error writing to file: ", err)
		file.Close()
		return
	}
	file.Close()
	fmt.Println("Content successfully written to file!")


}