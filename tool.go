package main

import (
  "os"
  "fmt"
  "github.com/gdamore/tcell/v2"
)

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

  style := tcell.StyleDefault.Foreground(tcell.ColorWhite)

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
        os.Exit(0)
      case tcell.KeyRune:
        switch ev.Rune() {
        case 'q', 'Q':
          s.Fini()
          os.Exit(0)
        }
      }
    case *tcell.EventMouse:
      x, y := ev.Position()
      s.Clear()
      s.SetContent(x, y, tcell.RuneBlock, nil, style)
      s.Sync()
    }
  }
}
