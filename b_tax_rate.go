package main

import "fmt"

func main() {
    var N int
    fmt.Scan(&N)
    a := false
    b := 0.0
    for i := 1; i <= N; i++ {
        b = 1.08 * float64(i)
        if int(b) == N {
            a = true
            fmt.Println(i)
            break
        }
    }
    if !a {
        fmt.Println(":(")
    }
}
