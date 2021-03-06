## Description

**asciibrot** is a simple ascii mandelbrot fractal generator running in terminal. 
It's written in Go and should run on all the existing platforms, however the Linux based is the one on which it was tested.

### Install

```
go get github.com/esimov/asciibrot
```

### Usage

```
go run julia.go --help
```

You can run the example in monochrome or color version.
For the color version use `--color` or `-c`. For monochrome version use `--mono` or `-m`.

You can build the binary version with `go build github.com/esimov/asciibrot`.

### Code example

To generate different output you can play with values defined in the main function:

```go
for {
    n += 0.045
    zoom += 0.04 * math.Sin(n)
    asciibrot.DrawFractal(zoom, math.Cos(n), math.Sin(n)/zoom*0.02, math.Sin(n), MAX_IT, true, isColor)

    // On CTRL+C restore default terminal foreground and background color
    go func() {
        <-c
        fmt.Fprint(asciibrot.Screen, "%s%s", "\x1b[49m", "\x1b[39m")
        fmt.Fprint(asciibrot.Screen, "\033[2J")
        asciibrot.Flush()
        os.Exit(1)
    }()
}
```

Blog post on my personal website: http://esimov.com/2016/05/ascii-mandelbrot-renderer-in-go

### Sample

![asciibrot-gif](https://user-images.githubusercontent.com/883386/68360648-9303ad00-0129-11ea-8db0-30a1f0dc9cb5.gif)

## License
This project is under MIT License.
