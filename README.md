# ll

short for "linelogger"

## [Read the Docs](https://godoc.org/github.com/jmervine/ll)

## Example

```go
package main
import (
    "gopkg.in/jmervine/ll.v1"
    "time"
)

func main() {
    begin := time.Now()

    // do stuff
    ll.Log(&begin, map[string]interface{}{
        "at": "main",
        "foo": "bar",
    })
}
```

*outputs*

```
YYYY-MM-DD HH:MM:SS at=main foo=bar duration=1.987us
```

