package chronos

import (
	"fmt"
	"strings"
	"time"
)

const _earliestSupportedYear = 1900
const _latestSupportedYear = 2100

type Lunar struct {
	time.Time
	year         int
	month        int
	day          int
	hour         int
	leapMonth    int
	leap         bool
	lichunOffset int //-1 when not meat lichun
}

var loc *time.Location

func init() {
	loc, _ = time.LoadLocation("Local")
}

func (lunar *Lunar) isLeap() bool {
	if lunar.leap && (lunar.month == lunar.leapMonth) {
		return true
	}
	return false
}

func (l *Lunar) Type() string {
	return "lunar"
}

func (l *Lunar) SetLichunOffset(fix int) {
	l.lichunOffset = fix
}

func (l *Lunar) Calendar() Calendar {
	t := time.Time{}
	t.AddDate(l.year, l.month, l.day)
	return New(t)
}

func (l *Lunar) EightChar() []string {
	rlt := l.YearString(l.lichunOffset) + l.MonthString() + l.DayString() + l.HourString()
	return strings.Split(rlt, "")
}

func (l *Lunar) HourString() string {
	return StemBranchHour(l.Year(), int(l.Month()), l.Day(), l.Hour())
}

func (l *Lunar) DayString() string {
	if l.Hour() >= 23 {
		return StemBranchDay(l.Year(), int(l.Month()), l.Day()+1)
	}
	return StemBranchDay(l.Year(), int(l.Month()), l.Day())
}

func (l *Lunar) MonthString() string {
	return StemBranchMonth(l.Year(), int(l.Month()), l.Day())
}

func (l *Lunar) YearString(fix int) string {
	if l.Month() > 2 || (l.Month() == 2 && l.Day() >= lichunDay(l.Year())) {
		return StemBranchYear(l.Year() + fix)
	}
	return StemBranchYear(l.Year() - 1)
}

func GetZodiac(l *Lunar) string {
	s := string([]rune(l.YearString(l.lichunOffset))[1])
	for idx, v := range earthyBranch {
		if strings.Compare(v, s) == 0 {
			return []string{
				`鼠`, `牛`, `虎`, `兔`, `龍`, `蛇`, `馬`, `羊`, `猴`, `子`, `狗`, `豬`,
			}[idx]
		}
	}
	return ""
}

func yearDay(y int) int {
	i, sum := 348, 348
	for i = 0x8000; i > 0x8; i >>= 1 {
		if (GetLunarInfo(y) & i) != 0 {
			sum++
		}
	}
	return sum + leapDay(y)
}

func leapDay(y int) int {
	if leapMonth(y) != 0 {
		if (GetLunarInfo(y) & 0x10000) != 0 {
			return 30
		}
		return 29
	}
	return 0
}

func leapMonth(y int) int {
	return GetLunarInfo(y) & 0xf
}

func monthDays(y int, m int) int {
	if m > 12 || m < 1 {
		return -1
	}
	if GetLunarInfo(y)&(0x10000>>uint32(m)) != 0 {
		return 30
	}
	return 29
}

func solarDays(y, m int) int {
	if m > 12 || m < 1 {
		return -1
	}
	var idx = m - 1
	if idx == 1 { // Feb case
		if (y%4 == 0) && (y%100 != 0) || (y%400 == 0) {
			return 29
		}
		return 28
	}
	monthDay := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	return monthDay[idx]
}

func GetAstro(m, d int) string {
	arr := []int{20, 19, 21, 21, 21, 22, 23, 23, 23, 23, 22, 22}
	idx := d < arr[m-1]
	index := m * 2
	if idx {
		index = m*2 - 2
	}
	var constellation = []string{
		`魔羯`, `水瓶`, `雙魚`, `白羊`, `金牛`, `雙子`, `巨蟹`, `獅子`, `處女`, `天秤`, `天蠍`, `射手`,
	}
	return constellation[index] + "座"
}

func lunarYear(offset int) (i int, restOffset int) {
	for i = _earliestSupportedYear; i <= _latestSupportedYear; i++ {
		day := yearDay(i)
		if offset-day < 1 {
			break
		}
		offset -= day
	}
	return i, offset
}

func lunarStart() time.Time {
	loc, _ := time.LoadLocation("Local")
	start, err := time.ParseInLocation("2006/01/02", "1900/01/30", loc)
	if err != nil {
		fmt.Println(err.Error())
	}
	return start
}

func lunarInput(date string) time.Time {
	input, err := time.ParseInLocation(_dateFormat, date, loc)
	if err != nil {
		fmt.Println(err.Error())
		return time.Time{}
	}
	return input
}

func CalculateLunar(date string) *Lunar {
	input := lunarInput(date)
	lunar := Lunar{
		Time: input,
		leap: false,
	}

	i, day := 0, 0
	isLeapYear := false

	start := lunarStart()
	offset := daysBetween(input, start)
	year, offset := lunarYear(offset)
	lunar.leapMonth = leapMonth(year) //計算該年閏哪個月

	//設定當年是否有閏月
	if lunar.leapMonth > 0 {
		isLeapYear = true
	}

	for i = 1; i <= 12; i++ {
		if i == lunar.leapMonth+1 && isLeapYear {
			day = leapDay(year)
			isLeapYear = false
			lunar.leap = true
			i--
		} else {
			day = monthDays(year, i)
		}
		offset -= day
		if offset <= 0 {
			break
		}
	}

	offset += day
	lunar.month = i
	lunar.day = offset
	lunar.year = year
	return &lunar

}

func daysBetween(d time.Time, s time.Time) int {
	newInput, err := time.ParseInLocation(_lunaDateFormat, d.Format(_lunaDateFormat), loc)
	if err != nil {
		return 0
	}

	subValue := float64(newInput.Unix()-s.Unix())/86400.0 + 0.5
	return int(subValue)
}

func LunarString(time time.Time) string {
	lunar := CalculateLunar(time.Format(_dateFormat))
	result := StemBranchYear(lunar.year) + "年"
	if lunar.leap && (lunar.month == lunar.leapMonth) {
		result += "閏"
	}
	result += getChineseMonth(lunar.month)
	result += getChineseDay(lunar.day)
	return result
}

func (lunar *Lunar) String() string {
	result := getChineseYear(lunar.year)
	if lunar.isLeap() {
		result += "閏"
	}
	result += getChineseMonth(lunar.month)
	result += getChineseDay(lunar.day)
	return result
}

var solarTerms = []string{
	`小寒`, `大寒`, `立春`, `雨水`, `驚蟄`, `春分`, `清明`, `穀雨`, `立夏`, `小滿`, `芒種`, `夏至`, `小暑`, `大暑`, `立秋`, `處暑`, `白露`, `秋分`, `寒露`, `霜降`, `立冬`, `小雪`, `大雪`, `冬至`,
}
