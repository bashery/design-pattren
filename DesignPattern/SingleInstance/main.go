package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var once sync.Once

func main() {
    go func() {
        time.Sleep(time.Microsecond*1010)
        //singleInstance = nil
    }()

    for i := 0; i< 50; i++ {
        wg.Add(1)
        go getInstance()
        time.Sleep(time.Second/5)
    }
    wg.Wait()
    fmt.Println("test")
}


type single struct{}

var singleInstance *single

func getInstance() *single {
    defer wg.Done()
    
    if singleInstance == nil {
        once.Do(
            func() {
                fmt.Println("Createing Single Instance Now")
                singleInstance = &single{}
            },
        )
    } else {
        fmt.Println("Single Instance alredy created")
    }
    return singleInstance
}

