package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputAnswerFromStdin(questions, inputs []string) error {
	reader := bufio.NewReader(os.Stdin)
	var err error
	for i:=0; i<len(questions); i++ {
		fmt.Println(questions[i])
		inputs[i], err = reader.ReadString('\n')
		if err!=nil {
			return err
		}
	}
	return err
}

func writeQuesAnsToFile(questions, inputs []string, file *os.File) error {
	var s string = ""
	for i:=0; i<len(questions); i++ {
		s = s + questions[i] + "\n" + inputs[i] + "\n\n"
	}

	_, err := file.WriteString(s)
	return err
}

func main() {
	questions := []string{"What's your name?", "When is your birthday?", "What do you do?"}
	var answers []string = make([]string, 3)
	
	err := inputAnswerFromStdin(questions, answers)
	if err!=nil {
		fmt.Println("Error input from Stdin: ", err)
	}

	var file *os.File
	file, err = os.Create("Example.txt")
	if err!=nil {
		fmt.Println("Error create file: ", err)
		file.Close()
		return
	}

	err = writeQuesAnsToFile(questions, answers, file)
	if err!=nil {
		fmt.Println("Error writing to file: ", err)
		file.Close()
		return
	}
	file.Close()
	fmt.Println("Success input and writing to file")
}