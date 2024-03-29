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

// Gets file name to save as. 
func getFileName(s tcell.Screen, style tcell.Style) string {
  _, y := s.Size()

  writeToScreen(s, style, 0, y-1, "File name to save?: ")
  s.ShowCursor(20, y-1)
  s.Show()
  var fileName []rune
  for {
    switch ev := s.PollEvent().(type){
      case *tcell.EventKey:
      switch ev.Key() {
      case tcell.KeyEscape:
        s.Fini()
        os.Exit(0)
      case tcell.KeyEnter:
        writeToScreen(s, style, 1, y-4, string(fileName))
        s.Show()
        return string(fileName)
      case tcell.KeyRune:
        fileName = append(fileName, ev.Rune())
        writeToScreen(s, style, 20, y-1, string(fileName))
        s.ShowCursor(20 + len(fileName), y-1)
        s.Show()
      case tcell.KeyBackspace, tcell.KeyBackspace2:
        tempLength := len(fileName)
        if len(fileName) > 0 {
          fileName = fileName[:len(fileName)-1]
        }
        difference := tempLength - len(fileName)
        // This is a hacky line for when you backspace, it covers the previous characters with a space.
        writeToScreen(s, style, 20, y-1, fmt.Sprintf(string(fileName)+strings.Repeat(" ", difference)))
        s.ShowCursor(20 + len(fileName), y-1)
        s.Show()
      }
    }
  }
}

// Writes the file, using the data entered and the name entered from getFileName()
func save(data [][]int, name string) {
  f, err := os.Create(fmt.Sprintf(name + ".txt"))
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

  _, y := s.Size()
  style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
  style2 := tcell.StyleDefault.Foreground(tcell.ColorGreen)

  s.SetStyle(style)
  s.EnableMouse()

  s.Clear()
  drawSlice(s, style2, data)
  writeToScreen(s, style, 0, y-1, "s: save | q: quit")
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
          //save(data, "gameoflife")
          fileName := getFileName(s, style)
          save(data, fileName)
          s.Fini()
          os.Exit(0)

        }
      }
    case *tcell.EventMouse:
      xPos, yPos := ev.Position()
      s.Clear()
      s.SetContent(xPos, yPos, tcell.RuneBlock, nil, style)
      drawSlice(s, style2, data)
      writeToScreen(s, style, 0, y-1, "s: save | q: quit")
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
