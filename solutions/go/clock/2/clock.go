package clock

import "fmt"

type Clock struct {
	hour, minute int
}

func New(h, m int) Clock {
	if m >= 60 {
		for m >= 60 {
			m -= 60
			h += 1
		}
	}
	if h >= 24 {
		h = h % 24
	}

	if m < 0 {
		for m < 0 {
			m += 60
			h -= 1
		}
	}

	if h < 0 {
		for h < 0 {
			h += 24
		}
	}

	return Clock{hour: h, minute: m}
}

func (c Clock) Add(m int) Clock {
	newHour := c.hour
	newMinute := c.minute

	newMinute += m
	if newMinute > 60 {
		newHour += 1
		newMinute -= 60
	}

	return New(newHour, newMinute)
}

func (c Clock) Subtract(m int) Clock {
	newHour := c.hour
	newMinute := c.minute

	newMinute -= m
	if newMinute < 0 {
		newHour -= 1
		newMinute += 60
	}

	return New(newHour, newMinute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
