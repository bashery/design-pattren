package main

import "fmt"

func main() {
    ada47, err := getGun("aka47")
    if err != nil {
        fmt.Println("error at nil pointer")
    }
    maverick, _ := getGun("maverick")
    printDetails(ada47)
    printDetails(maverick)
}

func printDetails(g iGun) {
    fmt.Printf("Gun: %s\n", g.getName())
    fmt.Printf("Gun: %d\n", g.getPower())
}
