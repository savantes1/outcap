# Standard Output Capture (outcap)

Simple library to catch stdout/stderr in Go. Cloned from https://github.com/PumpkinSeed/cage

#### Usage

```go
package main

import (
    "fmt"
    "os"

    "github.com/savantes1/outcap"
)

func main() {
    oc := outcap.NewContainer('\n')
    
    fmt.Println("test")
    fmt.Println("test2")
    fmt.Fprintln(os.Stderr, "stderr error")
    
    oc.Stop()
    fmt.Println(oc.Data)
    // [test, test2, stderr error]
}
```
