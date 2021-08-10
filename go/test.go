package main

import (
	"fmt"
	"math"
)

type Clock struct {
	h, m int
}

func New(h, m int) Clock {

	var tohour float64
	tohour = float64(m / 60)
	for m < 0 {
		m += 60
	}
	fmt.Println("ok", m)

	h += int(tohour)
	h = h % 24
	for h < 0 {
		h += 24
	}

	c := Clock{h, int(math.Abs(float64(m)))}
	return c
}

func main() {
	fmt.Println(New(-3, -70).String())
}

func (c Clock) String() string {
	houre := fmt.Sprintf("%d", c.h)
	minute := fmt.Sprintf("%d", c.m)
	if len(houre) < 2 {
		houre = "0" + houre
	}
	if len(minute) < 2 {
		minute = "0" + minute
	}
	return houre + ":" + minute
}

func (c Clock) Add() Clock {
	return c
}
