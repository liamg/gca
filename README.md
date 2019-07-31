# Go Cellular Automata

[![GoDoc](https://godoc.org/github.com/liamg/gca?status.svg)](https://godoc.org/github.com/liamg/gca)

A simple cellular automata package, useful for procedural generation.

Configurable, but has a helpful out of the box config that meets basic needs. Currently uses [Moore neighbourhoods](https://en.wikipedia.org/wiki/Moore_neighborhood).

![Demo](demo.gif)

You can see the code used to generate this gif in [examples/gif/main.go](examples/gif/main.go).

## Installation

GCA supports go modules, but will also work fine without, just use:

```
go get -u github.com/liamg/gca
```

## Usage Example

```
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

```

...will output something like:

```
#       ##     ####      ###### #######     ###  ####    ##  #######         #  
        #                 ####   ### ##           ##    #    ######         ##  
#     ####           ##    ##   ###    ##              ###   ######         ##  
#     ## ##      #   ###       ####     ##             ## #  ######         ##  
#                #    ###     #####                       ## #####          ####
                       ##     #####                       ##  ##             #  
 #          #                  ####                                             
### ##    ####                  ##                               ###            
#######   ###                                ###                 ###           #
#   ##    ##      #      #                  ####                 ###           #
                 ##     ##                 ###                  ###            #
                 #   ## ### ###           ###                  ####            #
       ##            ##  ######          #####                  ### ##         #
       ###    ##     ########           #######                 #######        #
       ## #  ####   ## #######    ##    ### ##       ##         ########  ##    
#        # ## ###  ##   #######   ###   ##          ###         ########  ###   
##      ###  ###  ##     #######   ##    ##         ##               ##    #    
###    ###    #   ##        #####  ##           ## ##       ##                 #
##    ####        ###        ####  ###          # ###                #         #
#      # ##        ##          #   ####         ## #                ###        #
#        ###        ###      ##     ####                            ###     ##  
#       ####        ###     ###      ####                            ##    ###  
#       ######     #####   ####     ######    ##     #               ##     #  #
#   ## ########   #############    #######    ###   ###          ### ###  ######
```
