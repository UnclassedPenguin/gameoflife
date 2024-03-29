# UnclassedPenguin Game of Life in Go

## Reasoning

I have been practicing my go lately, and so I decided to make [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life).  

Thanks to [The Coding Train](https://www.youtube.com/watch?v=FWSR_7kZuYg) for the complicated (at least for me) math to do the wrap around world.  

## Try it Out

To try it out yourself, make sure your go environment is set up, then:

```shell
$ go install github.com/unclassedpenguin/gameoflife@latest
```

Then simply run:

```shell
$ gameoflife
```

## Create Custom arrays

I created the tool.go to create custom starts. run tool.go `$ go run tool.go`. it starts with a blank screen. click on cells to enable them, or click again to disable them. When finished, hit s to save. It will ask for a file name, and then hit enter, and it will save filename with a .txt extension in the current folder. 
To run it, run the main program with -f and the file name. `$ go run main.go -f gameoflife.txt` then press 2. Press 2 at anytime to restart.  

## To-Do:

- Combine tool.go into main program.
- tool.go - Add a help pop up when press h.
- Handle panic if while running terminal size changes and it tries to go out of range of the array.
- Maybe add a history of what the starting array was, so you can repeat it?
- ~~tool.go - Add ability on save to enter a name of file to save.~~
  - This is kind of clunky. Its really cluttered and not very elegant. Can it be refactored?
- ~~Add ability to edit starts with the tool. So load an already created file into the tool, and edit it...Simple enough, ya?~~
- ~~Add ability to start with a custom array instead of random...So you can draw interesting patterns and see them working.~~
- ~~Need to add flags.~~ 
  - ~~One for a file to read, if you are going to do custom ones. Which will be an array of arrays.~~
  - ~~One to show generation.~~
