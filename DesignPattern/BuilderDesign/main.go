// Builder Pattern is a creational design pattern used for constructing complex objects.
package main

import "fmt"

func main() {
    normalBuilder := getBuilder("normallll")
    iglooBuilder := getBuilder("igloo")

    director := newDirector(normalBuilder)
    normalHouse := director.buildHouse()

    fmt.Printf("Normal House Door Type : %s\n", normalHouse.doorType)
    fmt.Printf("Normal House Window Type : %s\n", normalHouse.windowType)
    fmt.Printf("Normal House Num Floor : %d\n", normalHouse.floor)

    director.setBuilder(iglooBuilder)
    iglooHouse := director.buildHouse()

    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
    fmt.Printf("Igloo House window Type: %s\n", iglooHouse.windowType)
    fmt.Printf("Igloo House Door Type: %s\n", iglooHouse.doorType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}
