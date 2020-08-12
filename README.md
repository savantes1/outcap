# Standard Output Capture (outcap)

Simple library to catch stdout/stderr in Go. Originally cloned from https://github.com/PumpkinSeed/cage

#### Usage

```go
package main

import (
    "fmt"
    "os"

    "github.com/savantes1/outcap"
)

func main() {
    oc, err := outcap.NewContainer('\n')

    if err == nil {
    
        fmt.Println("test")
        fmt.Println("test2")
        fmt.Fprintln(os.Stderr, "stderr error")
        
        oc.Stop()
        fmt.Println(oc.OutData) // [test, test2]
        fmt.Println(oc.ErrorData) // [stderr error]

    } else {

        fmt.Fprintln(os.Stderr, err)
    }
}
```
