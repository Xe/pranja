package main

import (
	"log"

	perlin "github.com/aquilax/go-perlin"
	"github.com/bcicen/termui"
)

type Grid struct {
	cells [10][10]bool
	t     *termui.Table

	multiplier float64
	score      float64
	p          *perlin.Perlin
}

func (g *Grid) count() int {
	var result int
	for x := range g.cells {
		for y := range g.cells[x] {
			if g.cells[x][y] {
				result++
			}
		}
	}

	return result
}

func (g *Grid) spread(x, y int) {
	if g.cells[x][y] {
		g.score += 10
		g.cells[x][y] = false
		g.t.Rows[x][y] = doesntHaveDensity
		termui.Render(g.t)
	}

	if g.count() < 1 {
		g.fill()
	}
}

func (g *Grid) fill() {
	for x := range g.cells {
		for y := range g.cells[x] {
			val := g.p.Noise2D(float64(x)/10, float64(y)/10)

			if val > 0.1 {
				g.cells[x][y] = true
				g.t.Rows[x][y] = hasDensity
			} else {
				if len(g.cells[x]) != len(g.t.Rows[x]) {
					log.Printf("len(g.t.Rows[x]) = %v", len(g.t.Rows[x]))
					log.Fatalf("len(g.cells[x]) = %v", len(g.cells[x]))
				}

				g.cells[x][y] = false
				g.t.Rows[x][y] = doesntHaveDensity
			}
		}
	}

	termui.Render(g.t)
}
