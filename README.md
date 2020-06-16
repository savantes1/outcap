# Stdout Capture

Simple library to catch stdout/stderr in Go. Cloned from https://github.com/PumpkinSeed/cage

#### Usage

```go
package main

import (
    "fmt"
    "os"

    "github.com/savantes1/StdoutCapture"
)

func main() {
    c := cage.Start()
    
    fmt.Println("test")
    fmt.Println("test2")
    fmt.Fprintln(os.Stderr, "stderr error")
    
    cage.Stop(c)
    fmt.Println(c.Data)
    // [test, test2, stderr error]
}
```
