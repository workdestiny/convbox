# Convbox

[![Go Report Card](http://github.com/workdestiny/convbox)](http://github.com/workdestiny/convbox)

Convbox use convert number (integer) to short number (string)

## How to use
```go
func main() {
  
  // s string = "1.2K"
  s := convbox.ShortNumber(12300)

  // s string = "1.2M"
  s := convbox.ShortNumber(1230000)

}
```

## License
MIT