package main

import "fmt"

// button.go
type button struct {
    command command
}

func (b *button) press() {
    b.command.execute()
}

// command.go
type command interface {
    execute()
}

// onCommand.go
type onCommand struct {
    device device
}

func (c *onCommand) execute() {
    c.device.on()
}

// offCommand.go
type offCommand struct {
    device device
}

func (c *offCommand) execute() {
    c.device.off()
}

// device.go
type device interface {
    on()
    off()
}

// tv.go

type tv struct {
    isRunning bool
}

func (t *tv) on() {
    t.isRunning = true
    fmt.Println("Turning tv on")
}

func (t *tv) off() {
    t.isRunning = false
    fmt.Println("Turning tv off")
}

// main.go
func main() {
    tv := &tv{}
    onCommand := &onCommand{
        device: tv,
    }

    offCommand:= &offCommand{
        device:tv,
    }

    onButton := &button{
        command: onCommand,
    }
    onButton.press()
    offButton := &button{
        command:offCommand,
    }
    offButton.press()

}
