package main

import (
"fmt"
"github.com/cmbernard333/stringutil" // grab the package from directory $GOPATH/src/github.com/cmbernard333/stringutil
)

func main() {
    s := "Racecar"
    fmt.Printf("Hello, world!\n")
    fmt.Printf("%s\n",stringutil.Reverse(s)); // Don't need to quality stringutil here since its the only one called that
}
