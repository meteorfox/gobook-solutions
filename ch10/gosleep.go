package main

import (
        "fmt"
        "time"
)

func sleep(d time.Duration) {
        c := time.After(d)
        <-c
}

func main() {
        sleep(5 * time.Second)
        fmt.Println("Slept 5 seconds")
}
