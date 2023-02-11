package main

import (
  "os"
  "fmt"
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

func printSlice(slice [][]int) {
  for i, _ := range slice{
    fmt.Println(slice[i])
  }
}

func writeToScreen(s tcell.Screen, style tcell.Style, x int, y int, str string) {
  for i, char := range str {
    s.SetContent(x+i, y, rune(char), nil, style)
  }
}

func main() {
  s, err := tcell.NewScreen()
  if err != nil {
    fmt.Println("Error in tcell.NewScreen:", err)
  }

  if err := s.Init(); err != nil {
    fmt.Println("Error initializing screen:", err)
    os.Exit(1)
  }


  slice := createEmptySlice(s)
  x, y := s.Size()
  style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
  style2 := tcell.StyleDefault.Foreground(tcell.ColorGreen)
  style3 := tcell.StyleDefault.Foreground(tcell.ColorRed)

  s.SetStyle(style)
  s.EnableMouse()

  s.Clear()

  for {
    switch ev := s.PollEvent().(type) {
    case *tcell.EventResize:
      s.Sync()
    case *tcell.EventKey:
      switch ev.Key() {
      case tcell.KeyCtrlC, tcell.KeyEscape:
        s.Fini()
        printSlice(slice)
        fmt.Printf("X: %d\nY: %d\n", x, y)
        os.Exit(0)
      case tcell.KeyRune:
        switch ev.Rune() {
        case 'q', 'Q':
          s.Fini()
          printSlice(slice)
          os.Exit(0)
        }
      }
    case *tcell.EventMouse:
      x, y := ev.Position()
      s.Clear()
      s.SetContent(x, y, tcell.RuneBlock, nil, style)
      s.Sync()
      switch ev.Buttons() {
      case tcell.Button1:
        s.SetContent(x, y, tcell.RuneBlock, nil, style2)
        s.Sync()
      case tcell.Button2:
        s.SetContent(x, y, tcell.RuneBlock, nil, style3)
        s.Sync()
      }
    }
  }
}
