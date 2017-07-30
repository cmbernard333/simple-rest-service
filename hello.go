package main

import (
"fmt"
"github.com/cmbernard333/stringutil" // grab the package from directory $GOPATH/src/github.com/cmbernard333/stringutil
)

func main() {
    s := "!oG ,olleH"
    fmt.Printf("Hello, world!\n")
    // Don't need to quality stringutil here since its the only one called that
    fmt.Printf("Reverse %s -> %s\n",s,stringutil.Reverse(s)); //Printf is fairly similar to the C printf with format specifiers
}
