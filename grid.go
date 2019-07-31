package gca

import (
	"math/rand"
	"sync"
	"time"
)

type Grid struct {
	w                     int
	h                     int
	seed                  int64
	initialisationChance  float64
	data                  map[int]map[int]bool
	mut                   sync.RWMutex
	minNeighboursToBirth  int
	minNeighboursToRemain int
}

func NewGrid(width int, height int) *Grid {
	return &Grid{
		w:                     width,
		h:                     height,
		initialisationChance:  0.45,
		seed:                  time.Now().UnixNano(),
		data:                  map[int]map[int]bool{},
		minNeighboursToBirth:  5,
		minNeighboursToRemain: 2,
	}
}

func (g *Grid) SetSeed(seed int64) {
	g.seed = seed
}

func (g *Grid) SetMinNeighboursToBirth(count int) {
	g.minNeighboursToBirth = count
}

func (g *Grid) SetMinNeighboursToRemain(count int) {
	g.minNeighboursToRemain = count
}

func (g *Grid) SetInitialisationChance(chance float64) {
	g.initialisationChance = chance
}

func (g *Grid) Initialise() {
	g.mut.Lock()
	defer g.mut.Unlock()
	r := rand.New(rand.NewSource(g.seed))
	for x := 0; x < g.w; x++ {
		g.data[x] = map[int]bool{}
		for y := 0; y < g.h; y++ {
			g.data[x][y] = r.Float64() <= g.initialisationChance
		}
	}
}

func (g *Grid) Size() (int, int) {
	return g.w, g.h
}

func (g *Grid) Run(steps int) {
	for i := 0; i < steps; i++ {
		g.Step()
	}
}

func (g *Grid) countNeighbours(x, y int) int {
	var count int
	w, h := g.Size()
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if i < 0 || i >= w || j < 0 || j >= h {
				continue
			}
			if g.data[i][j] {
				count++
			}
		}
	}
	return count
}

func (g *Grid) Step() {
	g.mut.Lock()
	defer g.mut.Unlock()
	newMap := map[int]map[int]bool{}
	for x := 0; x < g.w; x++ {
		newMap[x] = map[int]bool{}
		for y := 0; y < g.h; y++ {
			count := g.countNeighbours(x, y)
			if count < g.minNeighboursToRemain {
				newMap[x][y] = false
			} else if count >= g.minNeighboursToBirth {
				newMap[x][y] = true
			} else {
				newMap[x][y] = g.data[x][y]
			}
		}
	}
	g.data = newMap
}

func (g *Grid) Read(x, y int) bool {
	g.mut.RLock()
	defer g.mut.RUnlock()
	col, ok := g.data[x]
	if !ok {
		return false
	}
	return col[y]
}
