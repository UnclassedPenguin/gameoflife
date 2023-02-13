package main

import (
  "os"
  "fmt"
)

func main() {
  fmt.Println("I'm going to write an array to a file")
  arr := [][]int{{0,0,0},{0,0,0},{0,0,0}}
  //arr := []int{0,0,1,1,0,0}
  fmt.Println(arr)

  f, err := os.Create("TESTTEST.txt")
  if err != nil {
    fmt.Println("Err Creating file: ", err)
  }

  defer f.Close()

  for _, value := range arr {
    fmt.Fprintln(f, value)
  }

}
