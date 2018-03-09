package main

import "fmt"
        "io"
        "os"
        "bufio"

func main() {
    // open input file
    fi, err := os.Open("final.txt")
    if err != nil {
        panic(err)
    }
  
  
