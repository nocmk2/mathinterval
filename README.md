
# mathinterval

parse math interval with golang


## Installing

go get -u github.com/nocmk2/mathinterval

## how to use

``` golang
package main

import (
    "fmt"

    "github.com/nocmk2/mathinterval"
)

func main() {
    res := mathinterval.Get("[-inf,100)").Hit(50)
    fmt.Println(res)
}
```
