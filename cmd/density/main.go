package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	perlin "github.com/aquilax/go-perlin"
	ui "github.com/bcicen/termui"
)

const (
	hasDensity        = `💢`
	doesntHaveDensity = `✅`
)

var (
	perlinAlpha = flag.Float64("perlin-alpha", 2, "Perlin noise alpha")
	perlinBeta  = flag.Float64("perlin-beta", 2, "Perlin noise beta")
	perlinN     = flag.Int("perlin-n", 20, "Perlin noise n")
)

func main() {
	flag.Parse()

	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	p := ui.NewPar("last x: -1")
	p.Width = 50
	p.Height = 4
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "Density Removal"
	p.BorderFg = ui.ColorCyan

	sc := ui.NewPar("0")
	sc.BorderLabel = "Score"
	sc.Height = 4

	rand.Seed(time.Now().Unix())

	g := &Grid{
		t: ui.NewTable(),
		p: perlin.NewPerlin(*perlinAlpha, *perlinBeta, *perlinN, rand.Int63()),
	}
	g.t.Width = 21
	g.t.Height = 21
	g.t.Border = true
	g.t.X = 0
	g.t.Y = 0
	g.t.TextAlign = ui.AlignCenter

	g.t.Rows = make([][]string, 10)
	for x := range g.t.Rows {
		g.t.Rows[x] = make([]string, 10)
		for y := range g.t.Rows[x] {
			g.t.Rows[x][y] = doesntHaveDensity
		}
	}

	g.spread(0, 0)

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(4, 0, p),
			ui.NewCol(4, 0, sc),
		),
		ui.NewRow(
			ui.NewCol(8, 0, g.t),
		),
	)

	ui.Body.Align()

	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	var lastPressedX = -1
	inputHandler := func(comp int) func(ui.Event) {
		return func(e ui.Event) {
			if lastPressedX != -1 && comp != -1 {
				g.spread(lastPressedX, comp)
				sc.Text = fmt.Sprintf("%d", int64(g.score))
				ui.Render(sc)
				lastPressedX = -1
			} else {
				lastPressedX = comp
			}

			p.Text = fmt.Sprintf("last x: %v", lastPressedX)
			ui.Render(p)
		}
	}

	ui.Handle("/sys/kdb/space", inputHandler(-1))
	ui.Handle("/sys/kbd/1", inputHandler(0))
	ui.Handle("/sys/kbd/2", inputHandler(1))
	ui.Handle("/sys/kbd/3", inputHandler(2))
	ui.Handle("/sys/kbd/4", inputHandler(3))
	ui.Handle("/sys/kbd/5", inputHandler(4))
	ui.Handle("/sys/kbd/6", inputHandler(5))
	ui.Handle("/sys/kbd/7", inputHandler(6))
	ui.Handle("/sys/kbd/8", inputHandler(7))
	ui.Handle("/sys/kbd/9", inputHandler(8))
	ui.Handle("/sys/kbd/0", inputHandler(9))

	ui.Loop()
}
