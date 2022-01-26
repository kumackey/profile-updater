package main

import (
	"bufio"
	"fmt"
	"os"
)

const filenameReadMe = "README.md"

func main() {
	f, err := os.Open(filenameReadMe)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	f, err = os.Create(filenameReadMe)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	for _, line := range lines {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
}
