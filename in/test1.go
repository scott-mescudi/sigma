package main

import ( 
    "fmt"
)

func somefunc (a int, c bool) int {
    var b int = 10
    if c {
        return a + b
    } else {
        return a * b
    }
}

func main() {
    fmt.Println(somefunc(21, true))
}
