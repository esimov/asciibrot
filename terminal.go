package asciibrot

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/mattn/go-colorable"
)

type winsize struct {
	Row uint16
	Col uint16
}

// Screen buffer
// Do not write to buffer directly, use package Print, Printf, Println functions instead.
var Screen *bytes.Buffer = new(bytes.Buffer)
var output *bufio.Writer = bufio.NewWriter(colorable.NewColorableStdout())

func init() {
	// Clear console
	output.WriteString("\033[2J")
	// Remove blinking cursor
	output.WriteString("\033[?25l")
}

// Get console width
func Width() int {
	ws, err := getWinsize()

	if err != nil {
		return -1
	}

	return int(ws.Row)
}

// Get console height
func Height() int {
	ws, err := getWinsize()
	if err != nil {
		return -1
	}
	return int(ws.Col)
}

// Flush buffer and ensure that it will not overflow screen
func Flush() {
	for idx, str := range strings.Split(Screen.String(), "\n") {
		if idx > Height() {
			return
		}

		output.WriteString(str + "\n")
	}

	output.Flush()
	Screen.Reset()
}

// Move cursor to given position
func MoveCursor(x int, y int) {
	fmt.Fprintf(Screen, "\033[%d;%dH", x, y)
}
