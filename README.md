# Hello World Golang Package
---

## Example Usage

main.go

```go
package main
 
import (
    "fmt"
    "github.com/yangpeng-chn/helloworld"
)
 
func main() {
    exa := test.Example{
        Text: "World",
    }
    fmt.Printf("Hello %s\n", exa.Hello())
}
``` 

```
$ go run main.go
Hello World
```
