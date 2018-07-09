package clock

import "fmt"

type Clock struct {
	hours, minutes int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}

func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

func New(h, m int) Clock {
	if m < 0 {
		h -= -m/60 + 1
		m -= 60 * (m/60 - 1)
	}
	if h < 0 {
		h -= 24 * (h/24 - 1)
	}

	return Clock{hours: (h + m/60) % 24, minutes: m % 60}
}
