
# mathinterval

parse math interval with golang

Getting Started
===============

## Installing

go get -u github.com/nocmk2/mathinterval

## Get an expression and hitted by a value

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
