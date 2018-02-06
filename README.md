# Convbox

[![Go Report Card](https://goreportcard.com/badge/github.com/workdestiny/convbox)](https://goreportcard.com/report/github.com/workdestiny/convbox)

Convbox use convert number (integer) to short number (string)

## How to use
```go
func main() {
  
  // s string = "12.3K"
  s := convbox.ShortNumber(12300)

  // s string = "1.2M"
  s := convbox.ShortNumber(1230000)

}
```

## License
MIT