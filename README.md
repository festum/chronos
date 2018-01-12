# Calendar date interface

```go
 chronos.New() // Create current date
 chronos.New(time.Now()) // Same as above
 chronos.New("2017/11/14 08:17") // Create the specified date
```

# Lunar calendar date

```
chronos.New().Lunar() //get the monthly calendar
```

# Solar calendar date

```go
chornos.New().Solar() //acquire calendar
```
