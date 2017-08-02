package main

import (
"fmt"
"github.com/cmbernard333/stringutil" // grab the package from directory $GOPATH/src/github.com/cmbernard333/stringutil
)

func forLoopTest () {
    // two ways to declare arrays
    // (1)
    // var x [4] int
    // x[n] = number
    //
    // (2)
    // x := [4]int{3,6,9,12} // initializer for an array of four ints
    
    x := [4]int{3,6,9,12} // initializer for an array of four ints
    // the range function returns the index and the value at the index in an array
    for i, c := range x {
        fmt.Printf("index %d, value %d\n", i, c)
    }
}

func main() {
    s := "!oG ,olleH"
    fmt.Printf("Hello, world!\n")
    // Don't need to quality stringutil here since its the only one called that
    fmt.Printf("Reverse %s -> %s\n",s,stringutil.Reverse(s)); //Printf is fairly similar to the C printf with format specifiers

    forLoopTest()
}
