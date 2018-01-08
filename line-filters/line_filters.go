package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {
	file, _ := os.Open("/tmp/lines")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		upperCaseText := strings.ToUpper(scanner.Text())
		fmt.Println(upperCaseText)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
