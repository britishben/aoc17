package main

import(
    "fmt"
    "os"
    "bufio"
    "aoc17/day1/captcha"
)

func main(){
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Hello")
    fmt.Println("Enter code:")
    scanner.Scan()
    ans := scanner.Text()
    fmt.Printf("Answer one is: %d.\n", captcha.CodeOne(ans))
    fmt.Printf("Answer two is: %d.\n", captcha.CodeTwo(ans))
}

