package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func bufioReader() {
	input := strings.NewReader("this is string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}

func bufioWriter() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello world\n")
	_, _ = writer.WriteString("Learn Golang\n")
	writer.Flush()
}

func main() {
	bufioReader()
	bufioWriter()
}
