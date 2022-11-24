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
          arr := createRandomArr(s)
          mainLoop(arr, s, style)
        }
      }
    }
  }
}

func mainLoop(arr [][]cell, s tcell.Screen, style tcell.Style) {

  for i := 0; i < 1; i++ {
    draw(arr, s, style)
    arr = countNeighbors(arr, s)
    for _, innerArr := range arr {
        //fmt.Print(innerArr)
      for _, cell := range innerArr {
        //fmt.Println(cell)
        if cell.state {
          if cell.neighbors < 2 {
            //cell.state = false
            arr[cell.x][cell.y].state = false
          } else if cell.neighbors > 3 {
            //cell.state = false
            arr[cell.x][cell.y].state = false
          }
        } else {
          if cell.neighbors == 3 {
            //cell.state = true
            arr[cell.x][cell.y].state = true
          }
        }
      }
    }
    time.Sleep(time.Millisecond * 1000)
  }
}

func countNeighbors(arr [][]cell, s tcell.Screen) [][]cell{
//func countNeighbors(arr [][]cell, s tcell.Screen) {
  x, y := s.Size()

  //s.Fini()

  // This is complete gibberish. Oof.
  // It iterates over every space and checks the spots around it
  // and does neighbors++ if a cell.state is true. Ignores
  // the border cells for the moment because you get out of range
  // errors.
  for i := 0; i < x; i++ {
    for j := 0; j < y; j++ {
      //fmt.Println("array[i][j]:", i, j)
      //fmt.Println("arr.x", arr[i][j].x)
      //fmt.Println("arr.y", arr[i][j].y)
      //fmt.Println("Neighbors:", arr[i][j].neighbors)
      if i > 1 && i < x-1 && j > 1 && j < y-1 {
        if arr[i-1][j-1].state {
          arr[i][j].neighbors++
        }
        if arr[i][j-1].state {
          arr[i][j].neighbors++
        }
        if arr[i+1][j-1].state {
          arr[i][j].neighbors++
        }
        if arr[i-1][j].state {
          arr[i][j].neighbors++
        }
        if arr[i+1][j].state {
          arr[i][j].neighbors++
        }
        if arr[i-1][j+1].state {
          arr[i][j].neighbors++
        }
        if arr[i][j+1].state {
          arr[i][j].neighbors++
        }
        if arr[i+1][j+1].state {
          arr[i][j].neighbors++
        }
      }
    }
  }
  return arr
}

func createRandomArr(s tcell.Screen) [][]cell {
  x, y := s.Size()
  var arr [][]cell

  for i := 0; i < x; i++ {
    var newArr []cell
    arr = append(arr, newArr)
    for j := 0; j < y; j++ {
      var newCell cell
      newCell.x = i
      newCell.y = j
      newCell.state = flipCoin(50,10)
      arr[i] = append(arr[i], newCell)
    }
  }
  return arr
}

// Draws a 2d slice of cells to the screen
func draw(arr [][]cell, s tcell.Screen, style tcell.Style) {
  s.Clear()
  for _, innerArr := range arr {
    for _, cell := range innerArr {
      if cell.state {
        s.SetContent(cell.x, cell.y, tcell.RuneBlock, []rune{}, style)
      }
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
// If you use flipCoin(10,2) it picks a random number between 0 and 9 and
// returns true if the number is less than 2. This way you can change
// the probability that you will get a live cell. 
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
