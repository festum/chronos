# Chronos

A Chinese day converter

## Example

Calendar date interface

```go
chronos.New() // Base on current date
chronos.New(time.Now()) // Time interface
chronos.New("2017/11/14 08:17") // string interface
```

Lunar calendar date

```go
chronos.New().Lunar() //get monthly calendar
```

Solar calendar date

```go
chornos.New().Solar() //get calendar
```
