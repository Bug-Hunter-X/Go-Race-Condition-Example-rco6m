```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mu sync.Mutex // add mutex
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mu.Lock() // lock before accessing count
                        count++
                        mu.Unlock() // unlock after accessing count
                }()
        }
        wg.Wait()
        fmt.Println(count) // this should always print 1000
}
```