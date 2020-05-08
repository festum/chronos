package chronos

import (
	"time"
)

const (
	_dateFormat     = "2006/01/02 15:04"
	_lunaDateFormat = "2006/01/02"
)

type calendar struct {
	time time.Time
}

type Calendar interface {
	Lunar() *Lunar
	Solar() *Solar
	LunarDate() string
}

type CalendarData interface {
	Type() string
	Calendar() Calendar
}

//Input support three type of time to create the calendar
//"2006/01/02 03:04" format string
// time.Time value or nil to create a new time.Now() value
func New(v ...interface{}) (c Calendar) {
	if v == nil {
		return &calendar{time.Now()}
	}
	switch vv := v[0].(type) {
	case string:
		return formatDate(vv)
	case time.Time:
		return &calendar{vv}
	}
	return
}

func formatDate(s string) Calendar {
	t, err := time.Parse(_dateFormat, s)
	if err != nil {
		t = time.Now()
	}
	return &calendar{
		time: t,
	}
}

func (c *calendar) Lunar() *Lunar {
	return CalculateLunar(c.time.Format(_dateFormat))
}

func (c *calendar) Solar() *Solar {
	return &Solar{time: c.time}
}

func (c *calendar) LunarDate() string {
	return c.Lunar().Date()
}
