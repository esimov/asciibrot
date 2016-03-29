package asciibrot

import (
	"fmt"
	"time"
	"sync"
)

var (
	wg      sync.WaitGroup
	height	int = Height()
	width	int = Width()
)

func DrawFractal(zoom, moveX, moveY float64, z float64, max_it int, update bool, isColor bool) {
	ticker := time.Tick(time.Millisecond * 2)

	charTable := map[int]string{1: "˙", 2: "˚", 3: "+", 4: "$", 5: "%", 6: "^", 7: "*", 8: "'", 9: "`"}
	foregroundColors := map[int]string{1: "\x1b[40m", 2: "\x1b[100m", 3: "\x1b[41m", 4: "\x1b[101m", 5: "\x1b[44m", 6: "\x1b[43m", 7: "\x1b[43m", 8: "\x1b[103m", 9: "\x1b[101m"}

	for row := 0; row < width; row++ {
		wg.Add(1)
		<-ticker
		Flush()

		go func(row int) {
			defer wg.Done()
			for col := 0; col < height; col++ {
				MoveCursor(row, col)

				newRe := 1.5 * (float64(row) - float64(width) / 2.0) / (0.5 * zoom * float64(width)) + moveX
				newIm := 1.2 * (float64(col) - float64(height) / 2.0) / (0.5 * zoom * float64(height)) + moveY

				var i = iterator(newRe, newIm, z, max_it)

				if i < max_it {
					if i > 6 {
						if isColor {
							fmt.Fprintf(Screen, "π%s%s%s%s%s", "\x1b[39m", "\x1b[1m", "\x1b[49m", "\x1b[41;32m", "\x1b[0m")
						} else {
							fmt.Fprintf(Screen, "π")

						}
					} else {
						if _, ok := charTable[i]; ok {
							if isColor {
								fmt.Fprintf(Screen, "%s%s", charTable[i], foregroundColors[i])
							} else {
								fmt.Fprintf(Screen, "%s", charTable[i])
							}
						}
					}
				} else {
					fmt.Fprintf(Screen, "∞")

				}
			}
			if !update {
				fmt.Println()
			}
		}(row)
	}

	wg.Wait()
}

func iterator(cx, cy float64, z float64, maxIter int) int {
	var iteration int = 0

	for iteration < maxIter {
		var x, y float64 = cx, cy
		cx = x*x - y*y + -0.95
		cy = 2*x*y + z

		if cx*cx+cy*cy > 4 {
			break
		}
		iteration++
	}
	return iteration
}