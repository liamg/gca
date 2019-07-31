package main

import (
	"fmt"

	"github.com/liamg/gca"
)

func main() {

	grid := gca.NewGrid(80, 24)
	grid.Initialise()
	grid.SetMinNeighboursToBirth(5)
	grid.SetMinNeighboursToRemain(2)
	grid.Run(3)

	w, h := grid.Size()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid.Read(x, y) {
				fmt.Printf(" ")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}

}
