package main

import "fmt"

func main() {
    var n, a, b int
    fmt.Scan(&n, &a, &b)
    p := a + b
    var s string
    fmt.Scan(&s)

    for i := 0; i < len(s); i++ {
        if s[i] == 'a' {
            if p > 0 {
                fmt.Println("Yes")
                p--
            } else {
                fmt.Println("No")
            }
        } else if s[i] == 'b' {
            if p > 0 && b > 0 {
                fmt.Println("Yes")
                p--
                b--
            } else {
                fmt.Println("No")
            }
        } else if s[i] == 'c' {
            fmt.Println("No")
        }
    }
}
