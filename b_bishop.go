//Problem Statement
//We have a board with H horizontal rows and W vertical columns of squares. There is a bishop at the top-left square on this board. How many squares can this bishop reach by zero or more movements?

package main

import "fmt"

func main() {
    var N, M int
    fmt.Scan(&N, &M)

    A := N * M
    if N == 1 || M == 1 {
        fmt.Println(1)
    } else if A%2 == 0 {
        fmt.Println(A / 2)
    } else {
        fmt.Println((A / 2) + 1)
    }
}
