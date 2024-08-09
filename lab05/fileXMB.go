 package main
 
 import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strconv"
	"math/rand"
	"time"
 )

 func main() {

	var input string
	var err error
	var X int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Nhap 1 so nguyen duong X: ")
	input, err = reader.ReadString('\n')
	if err!= nil {
		fmt.Println("Error input from Stdin: ", err)
		return
	}
	input = input[:len(input)-1] // loai bo \n
	X, err = strconv.Atoi(input)
	if err!=nil {
		fmt.Println("Error conver input string to number: ", err)
		return
	}

	var file *os.File
	file, err = os.Create("Example.txt")
	if err!=nil {
		fmt.Println("Error creating file: ", err)
		file.Close()
		return
	}


  chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	source := rand.NewSource(time.Now().UnixNano())
  rng := rand.New(source)
	var byteTotal uint64 = uint64(X*1048576)
	var count uint64 = 0
	var _ int
  for  i := uint64(0); i<byteTotal; i++ {
		count++
		if count%257==0 {	
			_, err = file.Write([]byte{byte('\n')})
			if err!=nil {
				if err == io.ErrShortWrite {
					fmt.Println("Error: Short write")
			  } else if err == io.ErrUnexpectedEOF {
					fmt.Println("Error: Unexpected EOF")
			  } else if os.IsNotExist(err) {
					fmt.Println("Error: File or directory does not exist")
			  } else if os.IsPermission(err) {
					fmt.Println("Error: Permission denied")
			  } else {
					fmt.Println("Error writing to file:", err)
			  }
				file.Close()
				return
			}
		} else {
			randomIndex := rng.Intn(len(chars))
      randomChar := chars[randomIndex]
			_, err = file.Write([]byte{byte(randomChar)})
			if err!=nil {
				if err == io.ErrShortWrite {
					fmt.Println("Error: Short write")
			  } else if err == io.ErrUnexpectedEOF {
					fmt.Println("Error: Unexpected EOF")
			  } else if os.IsNotExist(err) {
					fmt.Println("Error: File or directory does not exist")
			  } else if os.IsPermission(err) {
					fmt.Println("Error: Permission denied")
			  } else {
					fmt.Println("Error writing to file:", err)
			  }
				file.Close()
				return
			}
		}
	}
	
	file.Close()
}