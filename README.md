# go-interpolate
A small tool for easy string interpolation in golang.

## Usage

``` go
import (
  "fmt"
  "github.com/henvo/go-interpolate"
)

func main() {
  m := map[string]interface{}{
    "name": "Bob",
    "message": "You've got mail!"
  }

  fmt.Println(interpolate.FromMap("Hello %{name}! %{message}"))
}
```

Should print:

> Hello Bob! You've got mail! 

