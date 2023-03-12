package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Masukkan Input:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	charCount := make(map[string]int)

	for i := 0; i <= len(input)-1; i++ {
		char := string(input[i])
		fmt.Println(char)
		charCount[char]++
	}
	fmt.Println(charCount)
}
