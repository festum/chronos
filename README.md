# Chronos

![GitHub Test](https://github.com/festum/chronos/workflows/Test/badge.svg)

A Chinese day converter

## Example

Calendar date interface

```go
chronos.New() // Base on current date
chronos.New(time.Now()) // Time interface
d := chronos.New("2019/10/30 01:30") // string interface
l := d.Lunar() // Get lunar calendar
fmt.Println(l.Date())
// print: "己亥年十月初三"
fmt.Println(l.EightCharacter())
// print: "己 亥 甲 戌 庚 子 丁 丑"
d.Solar() // Get solar calendar
```
