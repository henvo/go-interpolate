# go-interpolate
[![Go Reference](https://pkg.go.dev/badge/github.com/henvo/go-interpolate.svg)](https://pkg.go.dev/github.com/henvo/go-interpolate)
[![codecov](https://codecov.io/gh/henvo/go-interpolate/graph/badge.svg?token=7ROIVRP5C0)](https://codecov.io/gh/henvo/go-interpolate)

A small tool for easy string interpolation in golang.

## Usage

### Values from `map[string]interface{}`

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

Will print:

> Hello Bob! You've got mail! 

### Values from `url.Values`

``` go
import (
  "fmt"
  "net/url"
  "github.com/henvo/go-interpolate"
)

func main() {
  m := make(url.Values)
  m.Add("name", "Bob")
  m.Add("message", "You've got mail!")

  fmt.Println(interpolate.FromURLValues("Hello %{name}! %{message}"))
}
```

Will print:

> Hello Bob! You've got mail!
