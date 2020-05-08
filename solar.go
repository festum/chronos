package chronos

import "time"

type Solar struct {
	time time.Time
}

func NewSolar(calendar Calendar) *Solar {
	return calendar.Solar()
}

func (s *Solar) Time() time.Time {
	return s.time
}
