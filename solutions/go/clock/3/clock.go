package clock

import "fmt"

const minutesInDay = 1440

type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	total := (h*60 + m) % minutesInDay
	if total < 0 {
		total += minutesInDay
	}
	return Clock{minutes: total}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
