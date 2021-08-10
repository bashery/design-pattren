package main

type iglooBuilder struct {
    windowType string
    doorType string
    floor int
}

func newIglooBuilder() *iglooBuilder {
    return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
    b.windowType = "Snow Widow"
}

func (b *iglooBuilder) setDoorType() {
    b.doorType = "Snow door"
}

func (b *iglooBuilder) setNumFloor() {
    b.floor = 1
}

func (b *iglooBuilder) getHouse() house {
    return house {
        doorType: b.doorType,
        windowType: b.windowType,
        floor: b.floor,
    }
}
