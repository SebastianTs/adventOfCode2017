package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	steps := parseInput("./input")
	dist, max := distance(steps)
	fmt.Printf("The first distance is: %d\nThe max distance is: %d\n", dist, max)
}

func distance(steps []string) (dist, max int) {

	p, _ := plot.New()
	p.Title.Text = "Way plotted"
	p.Add(plotter.NewGrid())
	pts := make(plotter.XYZs, len(steps))

	var x, y, z int = 0, 0, 0
	for i, step := range steps {
		switch step {
		case "n":
			y++
			z--
		case "ne":
			x++
			z--
		case "se":
			x++
			y--
		case "s":
			y--
			z++
		case "sw":
			x--
			z++
		case "nw":
			x--
			y++
		default:
		}

		pts[i].X = float64(x)
		pts[i].Y = float64(y)
		pts[i].Z = float64(z)
		dist = (abs(x) + abs(y) + abs(z)) / 2
		if max < dist {
			max = dist
		}
	}
	l, _ := plotter.NewLine(pts)
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}
	p.Add(l)
	p.Save(4*vg.Inch, 4*vg.Inch, "points.png")

	return
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func parseInput(file string) []string {
	out := make([]string, 0)

	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer fileHandle.Close()

	s := bufio.NewScanner(fileHandle)

	for s.Scan() {
		line := s.Text()
		out = strings.Split(line, ",")
		break
	}
	return out
}
