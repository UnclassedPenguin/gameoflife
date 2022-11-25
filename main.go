//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------
//
// Tyler(UnclassedPenguin) Conway's Game of Life 2022
//
// Author: Tyler(UnclassedPenguin)
//    URL: https://unclassed.ca
// GitHub: https://github.com/UnclassedPenguin
//
//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------

package main

import (
  "os"
  "fmt"
  "time"
  "math/rand"
  "github.com/gdamore/tcell/v2"
)


/*
These two variables are used in the initial slice 
to randomly decide wether a cell will start as 
alive or dead.

Used in the func createRandomSlice()

It says pick a random number between 0 and randHigh.
If it is less than randLow, make the cell alive.
*/
var randHigh = 50
var randLow = 10


// Displays the "menu" at the start.
func menu(s tcell.Screen, style tcell.Style) {
  x, y := s.Size()
  strings := []string{ "Unclassed Penguin Game of Life",
                       "Press 1 to start from random seed",
                       "(You can also press 1 at any time",
                       "while it is running to restart",
                       "with a new seed.)",
                       "Esc or Ctrl-C to quit",
                     }

  // Write strings to screen.
  for i, str := range strings {
    writeToScreen(s,style,((x/2)-(len(str)/2)),y/3+(i*2),str)
  }

  // Keyboard handling. Keys to quit (Esc, Ctrl-c, q)
  // and the key to start the game (1)
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
        case 'q', 'Q':
          s.Fini()
          os.Exit(0)
        case '1':
          slice := createRandomSlice(s)
          mainLoop(slice, s, style)
        }
      }
    }
  }
}

// This is the main loop. It gets a random slice, draws it to the screen,
// Does the computing to figure out how many neighbors a cell has, creates
// the new slice, and then draws it and repeats.
func mainLoop(slice [][]int, s tcell.Screen, style tcell.Style) {
  x, y := s.Size()

  // Handles keyboard input in main loop. Mainly ctrl-c 
  // so You can actually quit. Although 1 will start with
  // a new seed as well. 
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
          case 'q', 'Q':
            s.Fini()
            os.Exit(0)
          case '1':
            slice = createRandomSlice(s)
          }
        }
      }
    }
  }()

  // This is the main "draw" loop. Takes a slice, calculates
  // a new slice, draws it to the screen, and repeats.
  for {
    newSlice := createEmptySlice(s)
    for i := 0; i < x; i++ {
      for j := 0; j < y; j++ {
        var neighbors int
        neighbors = countNeighbors(s, slice, i, j)

        if slice[i][j] == 1 && (neighbors < 2 || neighbors > 3) {
          newSlice[i][j] = 0
        } else if slice[i][j] == 0 && neighbors == 3 {
          newSlice[i][j] = 1
        } else {
          newSlice[i][j] = slice[i][j]
        }
      }
      // This line makes it act strange. Not quite conways game of life,
      // But interesting none the less. Uncomment this line, and make sure
      // to comment out "slice = newSlice" that follows. Then go to top
      // and change the global variable randLow to 1 or 2

      //slice[i] = newSlice[i]
    }

    // This line makes it actually act like conways game of life. 
    // Need to increase global variable randLow to 10 or so.
    // You can comment out this line and uncomment the line above
    // For a different effect.

    slice = newSlice

    draw(slice, s, style)
    time.Sleep(time.Millisecond * 100)
  }
}

// This is a function that takes a slice and an x y position within that slice,
// and returns the count of how many of its 8 neighbors are "alive".
// ("alive" == 1)
func countNeighbors(s tcell.Screen, slice [][]int, x, y int) int{
  cols, rows := s.Size()
  neighbors := 0

  for i := -1; i < 2; i++ {
    for j := -1; j < 2; j++ {
      // Thanks to The Coding Train (on youtube) for this 
      // "wrap around" formula.
      neighbors += slice[(x+i+cols)%cols][(y+j+rows)%rows]
    }
  }

  // Don't count yourself as a neighbor.
  if slice[x][y] == 1 {
    neighbors--
  }

  return neighbors
}

// This creates the first random slice. 
// To tweak it, mess with the global variables at the top to change
// the probability that a cell will be alive.
func createRandomSlice(s tcell.Screen) [][]int {
  x, y := s.Size()
  var slice [][]int

  for i := 0; i < x; i++ {
    var newSlice []int
    slice = append(slice, newSlice)
    for j := 0; j < y; j++ {
      var newInt int
      // This is where the global variables change things.
      // Just changes the probability of returning true. 
      if flipCoin(randHigh,randLow) {
        newInt = 1
      } else {
        newInt = 0
      }
      slice[i] = append(slice[i], newInt)
    }
  }
  return slice
}

// I use this function to create an empty "2d" slice that then is populated 
// with the "real" data when we count neighbors on the working slice.
// (By "empty" I mean filled entirely with 0's)
// The slice is the same size as the terminal window. 
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


// Draws a 2d slice of cells to the screen
// 1's are white and 0's are background
func draw(slice [][]int, s tcell.Screen, style tcell.Style) {
  x, y := s.Size()
  s.Clear()
  for i := 0; i < x; i++ {
    for j := 0; j < y; j++ {
      if slice[i][j] == 1 {
        s.SetContent(i, j, tcell.RuneBlock, []rune{}, style)
      }
    }
  }
  s.Sync()
}

// This is used just to write strings to the screen. Used in the "menu".
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
