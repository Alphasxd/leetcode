package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    ch := make(chan struct{})

    go func() {
        defer wg.Done()
        for i := 1; i <= 10; i += 2 {
            fmt.Println("odd:", i)
            ch <- struct{}{}
        }
    }()

    go func() {
        defer wg.Done()
        for i := 2; i <= 10; i += 2 {
            <-ch
            fmt.Println("even:", i)
        }
    }()

    wg.Wait()
}