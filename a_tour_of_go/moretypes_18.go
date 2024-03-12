package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	canvas := make([][]uint8, dy)
	
	for i := range canvas {
    	canvas[i] = make([]uint8, dx)
	}
	
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			t := x^y
      // t *= t
			canvas[x][y] = uint8(t)
		}
	}
	return canvas
}

func main() {
	pic.Show(Pic)
}
