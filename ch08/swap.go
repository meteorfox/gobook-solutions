package main

import "fmt"

func main() {
        x := 1
        y := 2
        fmt.Println("Before swap:", x, y)
        swap(&x, &y)
        fmt.Println("After swap:", x, y)
}

func swap(x *int, y *int) {
        *x, *y = *y, *x
}
