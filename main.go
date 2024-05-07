package main

import (
    "fmt"
    "math/rand"
    "sync"
)

type MAP struct {
    sync.Mutex
    goroutine map[int]int
}

func (mp *MAP) read(ky int) (int, bool) {
    mp.Lock()
    defer mp.Unlock()
    value, ok := mp.goroutine[ky]
    return value, ok
}

func (mp *MAP) write(ky, value int) {
    mp.Lock()
    defer mp.Unlock()
    mp.goroutine[ky] = value
}

func (mp *MAP) delete(ky int) {
    mp.Lock()
    defer mp.Unlock()
    delete(mp.goroutine, ky)
}

func main() {
    var wg sync.WaitGroup
    mp := &MAP{goroutine: make(map[int]int)}

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            mp.write(i, rand.Intn(100))
        }(i)
    }

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            value, ok := mp.read(i)
            if ok {
                fmt.Printf("Key: %d, Value: %d\n", i, value)
            } else {
                fmt.Printf("Key: %d, Value: mavjud emas\n", i)
            }
        }(i)
    }

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            mp.delete(i)
        }(i)
    }

    wg.Wait()
}
