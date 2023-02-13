package main

import (
  //"os"
  "io/ioutil"
  "fmt"
  "strings"
)

func main() {
  // Read file
  f, _ := ioutil.ReadFile("TESTTEST.txt")

  lines := strings.Split(string(f), "\n")

  var cleanLines []string
  var data [][]int

  // Remove the leading [ and the trailing ] of each line and remove
  // empty lines
  for _, line := range lines {
    newline1 := strings.Replace(line, "[", "", -1)
    newline2 := strings.Replace(newline1, "]", "", -1)
    newline3 := strings.Replace(newline2, " ", "", -1)
    if len(line) > 0 {
      cleanLines = append(cleanLines, newline3)
    }
  }

  // Create an array of arrays with the proper size for the data. Fill with 0's
  for i := 0; i < len(cleanLines); i++ {
    var newSlice []int
    data = append(data, newSlice)
    for j := 0; j < len(cleanLines[0]); j++ {
      newInt := 0
      data[i] = append(data[i], newInt)
    }
  }

  // Go over the string data, and convert it to the int data. Only have to update
  // the 1's
  for i, line := range cleanLines {
    for j, char := range line {
      if char == '1' {
        data[i][j] = 1
      }
    }
  }

  fmt.Println("------------------")
  fmt.Println(data)
}
