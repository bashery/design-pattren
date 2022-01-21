/// this is cleares then fucking pattren designe
package main

import "fmt"
type device interface {
    on()
    off()
}
type tv struct {
    isRuning bool
}
func (t *tv) on() {
    t.isRuning = true
    fmt.Println("Turning device on")
}
func (t *tv) off () {
    t.isRuning = false 
    fmt.Println("Turning device Off")
}
func main() {
    var device device
    device = &tv{}
    device.on()
    device.off()
}

