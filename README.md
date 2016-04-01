### Description

**ASCIIbrot** is a simple Mandelbrot fractal generator running in terminal. 
It's written in Go and it was designed to be ran on different platforms, however the linux based is the one on which it was tested.

The code is meant to be as clear as possible, so it should be self explanatory, however comments are provided on less obvious code parts.  

You can get the library with the following command: 

```
go get github.com/esimov/asciibrot
```

### Usage

To run the mandelbrot generator type:
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

### Samples

**Ascii Mandelbrot snpashot - color version**
![Screenshot_1](https://raw.githubusercontent.com/esimov/asciibrot/master/examples/screenshot_1.png)

**Ascii Mandelbrot snpashot - monochrome version**
![Screenshot_2](https://raw.githubusercontent.com/esimov/asciibrot/master/examples/screenshot_2.png)