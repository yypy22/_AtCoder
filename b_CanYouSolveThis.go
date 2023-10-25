package main

import (
    "fmt"
)

func main() {
    var N, M, C int
    fmt.Scan(&N, &M, &C)

    B := make([]int, M)
    for i := 0; i < M; i++ {
        fmt.Scan(&B[i])
    }

    ans := 0
    for i := 0; i < N; i++ {
        a := make([]int, M)
        for j := 0; j < M; j++ {
            fmt.Scan(&a[j])
        }
        x := 0
        for j := 0; j < M; j++ {
            x += a[j] * B[j]
        }
        if x+C > 0 {
            ans++
        }
    }
    fmt.Println(ans)
}
