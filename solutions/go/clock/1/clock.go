package clock

import "fmt"

type Clock struct {
	hour, minute int
}

func New(h, m int) Clock {
	return Clock{hour: h, minute: m}
}

func (c Clock) Add(m int) Clock {
	return Clock{hour: c.hour, minute: c.minute + m}
}

func (c Clock) Subtract(m int) Clock {
	return Clock{hour: c.hour, minute: c.minute - m}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
