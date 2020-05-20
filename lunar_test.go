package chronos_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/festum/chronos"
	"github.com/stretchr/testify/assert"
)

func TestGetDayString(t *testing.T) {
	assert.Equal(t, "620419", strconv.Itoa(0x97783))
	assert.Equal(t, "621520", strconv.Itoa(0x97bd0))
	assert.Equal(t, "621622", strconv.Itoa(0x97c36))
	assert.Equal(t, "723823", strconv.Itoa(0xb0b6f))
}
func TestGetZodiac(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("豬", chronos.GetZodiac(chronos.New("2020/01/24 18:40").Lunar()))
	assert.Equal("牛", chronos.GetZodiac(chronos.New("1985/12/13 18:40").Lunar()))
	assert.Equal("鼠", chronos.GetZodiac(chronos.New("2020/02/11 18:40").Lunar()))
	assert.Equal("牛", chronos.GetZodiac(chronos.New("2021/02/12 18:40").Lunar()))
}

func TestStemBranchHour(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("丁酉", chronos.StemBranchYear(2017))
	assert.Equal("庚辰", chronos.StemBranchHour(2018, 1, 13, 8))
	assert.Equal("乙巳", chronos.StemBranchDay(2017, 11, 14))
	assert.Equal("庚辰", chronos.StemBranchHour(2017, 11, 14, 8))
}

func TestNewLunar(t *testing.T) {
	assert.NotNil(t, chronos.New().Lunar().Date())
}

func TestCalculateLunar(t *testing.T) {
	assert := assert.New(t)
	// assert.Equal("庚子年四月十六日", )
	assert.Equal(strings.Split("己 亥 甲 戌 壬 寅 庚 子", " "), chronos.New("2019/10/31 23:13").Lunar().EightCharacter())
	assert.Equal(strings.Split("己 亥 甲 戌 壬 寅 庚 子", " "), chronos.New("2019/11/01 0:13").Lunar().EightCharacter())
	assert.Equal(strings.Split("己 亥 甲 戌 壬 寅 辛 丑", " "), chronos.New("2019/11/01 1:13").Lunar().EightCharacter())
	assert.Equal(strings.Split("己 亥 丁 醜 丁 丑 庚 子", " "), chronos.New("2020/02/03 23:13").Lunar().EightCharacter())
	assert.Equal(strings.Split("庚 子 戊 寅 丁 丑 辛 丑", " "), chronos.New("2020/02/04 1:13").Lunar().EightCharacter())
	assert.Equal(strings.Split("庚 子 戊 寅 丁 丑 己 酉", " "), chronos.New("2020/02/04 18:13").Lunar().EightCharacter())
	assert.Equal("己亥年臘月三十", chronos.New("2020/01/24 0:40").LunarDate())
	assert.Equal("庚子年正月初一", chronos.New("2020/01/25 0:40").LunarDate())
	assert.Equal(strings.Split("戊 辰 乙 醜 丁 卯 己 酉", " "), chronos.New("1989/01/07 18:40").Lunar().EightCharacter())
	assert.Equal("戊辰年十一月三十", chronos.New("1989/01/07 18:40").LunarDate())
	assert.Equal("戊辰年十一月三十", chronos.New("1989/01/07 0:40").LunarDate())
	assert.Equal("己亥年四月廿八", chronos.New("2019/06/01 0:40").LunarDate())
}

func TestLunarEightCharacter(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(strings.Split("己 亥 甲 戌 庚 子 丁 丑", " "), chronos.New("2019/10/30 01:30").Lunar().EightCharacter())
	expected := strings.Split("己 亥 甲 戌 辛 丑 戊 子", " ")
	assert.Equal(expected, chronos.New("2019/10/30 23:00").Lunar().EightCharacter())
	assert.Equal(expected, chronos.New("2019/10/31 00:30").Lunar().EightCharacter())
	assert.Equal("己亥年十月初三", chronos.New("2019/10/30 23:00").Lunar().Date())
}
