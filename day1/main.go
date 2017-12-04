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
    ans := captcha.Code(scanner.Text())
    fmt.Printf("Answer is: %d", ans)
}

