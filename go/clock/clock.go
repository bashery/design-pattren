package clock

import (
	"fmt"
)

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {

	for minute < 0 {
		minute += 60
		hour--
	}
	for minute >= 60 {
		minute -= 60
		hour++
	}

	hour = hour % 24

	for hour < 0 {
		hour += 24
	}

	clock := Clock{hour, minute}
	return clock
}

func (c Clock) Add(m int) Clock {

	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (clock Clock) String() string {
	hour := fmt.Sprintf("%02d", clock.hour)
	minute := fmt.Sprintf("%02d", clock.minute)
	return hour + ":" + minute
}
