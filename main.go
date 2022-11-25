package main

import (
  "os"
  "fmt"
  "time"
  "math/rand"
  "github.com/gdamore/tcell/v2"
)

// This says pick a number between 1 and 50.
// If it is less than 2, make the cell alive.
var randHigh = 50
var randLow = 10

// Displays the menu at the start.
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

// This is the main loop. It gets a random array, draws it to the screen,
// Does the computing to figure out how many neighbors a cell has, creates
// the new array, and then draws it and repeats.
func mainLoop(arr [][]int, s tcell.Screen, style tcell.Style) {
  x, y := s.Size()

  go func() {
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
            arr = createRandomArr(s)
          }

        }
      }
    }
  }()

  for {
    newArr := createEmptyArr(s)

    for i := 0; i < x; i++ {
      for j := 0; j < y; j++ {
        var neighbors int
        neighbors = countNeighbors(s, arr, i, j)

        if arr[i][j] == 1 && (neighbors < 2 || neighbors > 3) {
          newArr[i][j] = 0
        } else if arr[i][j] == 0 && neighbors == 3 {
          newArr[i][j] = 1
        } else {
          newArr[i][j] = arr[i][j]
        }
      }
      // This line makes it act strange. Not quite conways game of life,
      // But interesting none the less. Uncomment this line, and make sure
      // to comment out "arr = newArr" that follows. Then go to top
      // and change the global variable randLow to 1 or 2

      //arr[i] = newArr[i]
    }

    // This line makes it actually act like conways game of life. 
    // Need to increase global variable randLow to 10 or so.
    // You can comment out this line and uncomment the line above
    // For a different effect.

    arr = newArr

    draw(arr, s, style)
    time.Sleep(time.Millisecond * 100)
  }
}

// This is a function that takes an array, and an x y position within that array,
// and returns the count of how many of its 8 neighbors are "alive".
// ("alive" == 1)
func countNeighbors(s tcell.Screen, arr [][]int, x, y int) int{
  cols, rows := s.Size()
  neighbors := 0

  for i := -1; i < 2; i++ {
    for j := -1; j < 2; j++ {
      neighbors += arr[(x+i+cols)%cols][(y+j+rows)%rows]
    }
  }

  if arr[x][y] == 1 {
    neighbors--
  }
  return neighbors
}

// This creates the first random array. 
// To tweak it, mess with the global variables at the top to change
// the probability that a cell will be alive.
func createRandomArr(s tcell.Screen) [][]int {
  x, y := s.Size()
  var arr [][]int

  for i := 0; i < x; i++ {
    var newArr []int
    arr = append(arr, newArr)
    for j := 0; j < y; j++ {
      var newInt int
      if flipCoin(randHigh,randLow) {
        newInt = 1
      } else {
        newInt = 0
      }
      arr[i] = append(arr[i], newInt)
    }
  }
  return arr
}

// I use this function to create an empty array that then is populated 
// with the "real" data when we count neighbors on the working array.
func createEmptyArr(s tcell.Screen) [][]int {
  x, y := s.Size()
  var arr [][]int

  for i := 0; i < x; i++ {
    var newArr []int
    arr = append(arr, newArr)
    for j := 0; j < y; j++ {
      newInt := 0
      arr[i] = append(arr[i], newInt)
    }
  }
  return arr
}


// Draws a 2d slice of cells to the screen
// 1's are white and 0's are background
func draw(arr [][]int, s tcell.Screen, style tcell.Style) {
  x, y := s.Size()
  s.Clear()
  for i := 0; i < x; i++ {
    for j := 0; j < y; j++ {
      if arr[i][j] == 1 {
        s.SetContent(i, j, tcell.RuneBlock, []rune{}, style)
      }
    }
  }
  s.Sync()
}

// This is used just to write strings to the screen. Used in the main menu.
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

// Main function. Starts the tcell.Screen and passes to the menu. 
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
