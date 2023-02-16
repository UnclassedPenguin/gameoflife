package main

import (
  "os"
  "fmt"
  "flag"
  "strings"
  "io/ioutil"
  "github.com/gdamore/tcell/v2"
)

func createEmptySlice(s tcell.Screen) [][]int {
  x, y := s.Size()
  var slice [][]int

  for i := 0; i < x; i++ {
    var newSlice []int
    slice = append(slice, newSlice)
    for j := 0; j < y; j++ {
      newInt := 0
      slice[i] = append(slice[i], newInt)
    }
  }
  return slice
}

func drawSlice(s tcell.Screen, style tcell.Style, slice [][]int) {
  x, y := s.Size()
  for i := 0; i < x; i++ {
    for j := 0; j < y; j++ {
      if slice[i][j] == 1 {
        s.SetContent(i,j,tcell.RuneBlock, nil, style)
      }
    }
  }
}

func printSlice(slice [][]int) {
  for i, _ := range slice{
    fmt.Println(slice[i])
  }
}

func updateData(x, y int, data [][]int) {
  if data[x][y] == 0 {
    data[x][y] = 1
  } else {
    data[x][y] = 0
  }
}

func save(data [][]int) {
  f, err := os.Create("gameoflife.txt")
  if err != nil {
    fmt.Println("Err Creating file: ", err)
  }

  defer f.Close()

  for _, value := range data {
    fmt.Fprintln(f, value)
  }
}

// Reads a file as a start
func readFile(file string) [][]int {
  f, _ := ioutil.ReadFile(file)

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
  return data

}


func writeToScreen(s tcell.Screen, style tcell.Style, x int, y int, str string) {
  for i, char := range str {
    s.SetContent(x+i, y, rune(char), nil, style)
  }
}

func main() {

  var file string

  flag.StringVar(&file, "f", "", "File to read from")

  flag.Parse()


  s, err := tcell.NewScreen()
  if err != nil {
    fmt.Println("Error in tcell.NewScreen:", err)
  }

  if err := s.Init(); err != nil {
    fmt.Println("Error initializing screen:", err)
    os.Exit(1)
  }


  var data [][]int

  if file != "" {
    data = readFile(file)
  } else {
    data = createEmptySlice(s)
  }

  style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
  style2 := tcell.StyleDefault.Foreground(tcell.ColorGreen)

  s.SetStyle(style)
  s.EnableMouse()

  s.Clear()
  drawSlice(s, style2, data)
  s.Sync()

  for {
    switch ev := s.PollEvent().(type) {
    case *tcell.EventResize:
      s.Sync()
    case *tcell.EventKey:
      switch ev.Key() {
      case tcell.KeyCtrlC, tcell.KeyEscape:
        s.Fini()
        //printSlice(data)
        os.Exit(0)
      case tcell.KeyRune:
        switch ev.Rune() {
        case 'q', 'Q':
          s.Fini()
          //printSlice(data)
          os.Exit(0)
        case 's', 'S':
          save(data)
          s.Fini()
          os.Exit(0)
        }
      }
    case *tcell.EventMouse:
      xPos, yPos := ev.Position()
      s.Clear()
      s.SetContent(xPos, yPos, tcell.RuneBlock, nil, style)
      drawSlice(s, style2, data)
      s.Show()
      switch ev.Buttons() {
      case tcell.Button1:
        updateData(xPos, yPos, data)
        s.SetContent(xPos, yPos, tcell.RuneBlock, nil, style2)
        s.Show()
      }
    }
  }
}
