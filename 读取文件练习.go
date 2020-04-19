package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "/Users/chris/Documents/girl.ini"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	buf := bufio.NewReader(file)

	for {
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		fmt.Println(string(b))
	}

	defer file.Close()
}
