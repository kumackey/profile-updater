package main

import (
	"bufio"
	"fmt"
	"os"
)

const filenameReadMe = "README.md"
const beginLine = "<!-- begin line of qiita profile -->"
const endLine = "<!-- end line of qiita profile -->"

type line string

func main() {
	f, err := os.Open(filenameReadMe)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []line
	for scanner.Scan() {
		lines = append(lines, line(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	f, err = os.Create(filenameReadMe)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	lines = replaceLines(lines)
	for _, line := range lines {
		_, err := f.WriteString(string(line) + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func replaceLines(lines []line) []line {
	replacedLines := make([]line, 0, len(lines))
	writeMode := false
	for _, line := range lines {
		if line == endLine {
			replacedLines = append(replacedLines, "replaced line")
			replacedLines = append(replacedLines, line)
			writeMode = false
			continue
		}

		if writeMode {
			continue
		}

		replacedLines = append(replacedLines, line)

		if line == beginLine {
			writeMode = true
		}
	}

	return replacedLines
}
