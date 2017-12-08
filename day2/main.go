package main

import (
	"aoc17/day2/checksum"
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hello")
	//fmt.Println("Enter code:")
	scanner.Scan()
	ans := scanner.Text()
	fmt.Printf("Answer one is: %d.\n", captcha.CodeOne(ans))
	fmt.Printf("Answer two is: %d.\n", captcha.CodeTwo(ans))
}
