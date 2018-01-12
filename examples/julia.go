package main

import (
	"fmt"
	"github.com/esimov/asciibrot"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	MAX_IT int = 1000
)

var zoom float64 = 1.0
var isColor bool = false

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	if len(os.Args) > 1 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println(`Usage go run mandelbrot_cli.go [--]
			-c --color		generate ASCII mandelbrot in color
			-m --mono		generate ASCII mandelbrot in monochrome`)
			os.Exit(1)
		}
		if os.Args[1] == "--color" || os.Args[1] == "-c" {
			isColor = true
		} else if os.Args[1] == "--mono" || os.Args[1] == "-m" {
			isColor = false
		}
	}

	zoom = 1.2 + rand.Float64()*1.8
	asciibrot.MoveCursor(0, 0)

	var n float64 = 20
	for {
		n += 0.045
		zoom += 0.04 * math.Sin(n)
		asciibrot.DrawFractal(zoom, math.Cos(n), math.Sin(n)/zoom*0.02, math.Sin(n), MAX_IT, isColor)

		// On CTRL+C restore default terminal foreground and background color
		go func() {
			<-c
			fmt.Fprint(asciibrot.Screen, "%s%s", "\x1b[49m", "\x1b[39m")
			fmt.Fprint(asciibrot.Screen, "\033[2J")
			asciibrot.Flush()
			os.Exit(1)
		}()
	}
}
