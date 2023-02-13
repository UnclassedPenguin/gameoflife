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


  data := createEmptySlice(s)
  //x, y := s.Size()
  style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
  style2 := tcell.StyleDefault.Foreground(tcell.ColorGreen)
  style3 := tcell.StyleDefault.Foreground(tcell.ColorRed)

  s.SetStyle(style)
  s.EnableMouse()

  s.Clear()

  for {
    s.Clear()
    switch ev := s.PollEvent().(type) {
    case *tcell.EventResize:
      s.Sync()
    case *tcell.EventKey:
      switch ev.Key() {
      case tcell.KeyCtrlC, tcell.KeyEscape:
        s.Fini()
        printSlice(data)
        os.Exit(0)
      case tcell.KeyRune:
        switch ev.Rune() {
        case 'q', 'Q':
          s.Fini()
          printSlice(data)
          os.Exit(0)
        }
      }
    case *tcell.EventMouse:
      xPos, yPos := ev.Position()
      s.SetContent(xPos, yPos, tcell.RuneBlock, nil, style)

      drawSlice(s, style3, data)
      s.Show()
      switch ev.Buttons() {
      case tcell.Button1:
        updateData(xPos,yPos, data)
        s.SetContent(xPos, yPos, tcell.RuneBlock, nil, style2)
        s.Show()
      case tcell.Button2:
        s.SetContent(xPos, yPos, tcell.RuneBlock, nil, style3)
        s.Show()
      }
    }
  }
}
