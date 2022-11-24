package main

import (
  "os"
  "fmt"
  "time"
  "math/rand"
  "github.com/gdamore/tcell/v2"
)

type cell struct {
  x int
  y int
  state bool
  neighbors int
}

func menu(s tcell.Screen, style tcell.Style) {
  x, y := s.Size()
  str1 := "Unclassed Penguin Game of Life"
  str2 := "Press 1 to start from random seed"
  str3 := "Esc or Ctrl-C to quit"

  writeToScreen(s,style,((x/2)-(len(str1)/2)),y/3,str1)
  writeToScreen(s,style,((x/2)-(len(str2)/2)),y/3+2,str2)
  writeToScreen(s,style,((x/2)-(len(str3)/2)),y/3+4,str3)

  for {
    switch ev := s.PollEvent().(type) {
    case *tcell.EventResize:
      s.Sync()
    case *tcell.EventKey:
      switch ev.Key() {
      case tcell.KeyCtrlC, tcell.KeyEscape:
        s.Fini()
        os.Exit(0)
      case tcell.KeyRune:
        switch ev.Rune() {
        case '1':
          arr := creatRandomArr(s)
          draw(arr, s, style)
        }
      }
    }
  }
}

//func countNeighors(arr []cell) []cell{
  //for _, cell := range arr {
    //for x := -1; x < 2; x++ {
      //for y := -1; y <2; y++ {
        //if arr[cell
  //}
//}

// This creates a random array that matches the size of the screen
func creatRandomArr(s tcell.Screen) []cell {
  x, y := s.Size()

  var arr []cell
  for i := 0; i < x; i++ {
    for j := 0; j < y; j++ {
      var newCell cell
      newCell.x = i
      newCell.y = j
      arr = append(arr, newCell)
    }
  }

  for i, _ := range arr {
    if flipCoin(50, 5) {
      arr[i].state = true
    }
  }

  return arr
}


// Draws a slice of cells to the screen
func draw(arr []cell, s tcell.Screen, style tcell.Style) {
  s.Clear()
  for _, cell := range arr {
    if cell.state {
      s.SetContent(cell.x, cell.y, tcell.RuneBlock, []rune{}, style)
    }
  }
  s.Sync()
}

func writeToScreen(s tcell.Screen, style tcell.Style, x int, y int, str string) {
  for i, char := range str {
    s.SetContent(x+i, y, rune(char), []rune{}, style)
  }
}

// I use this to decide if a cell has a state of true (alive) or state of
// false (dead) to start
func flipCoin(total int, limit int) bool{
  rand.Seed(time.Now().UnixNano())
  x := rand.Intn(total)
  if x <= limit {
    return true
  } else {
    return false
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

  s.Clear()

  style := tcell.StyleDefault.Foreground(tcell.ColorWhite)

  menu(s, style)
}
