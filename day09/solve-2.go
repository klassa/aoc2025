package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func ToPoint(rawCoords []string) *Point {
	x, err := strconv.Atoi(rawCoords[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(rawCoords[1])
	if err != nil {
		panic(err)
	}
	return &Point{x, y}
}

type Grid [][]byte

func (g *Grid) Init(maxX, maxY int) {
	for y := 0; y <= maxY+1; y++ {
		*g = append(*g, make([]byte, maxX+1))
	}
}

func (g *Grid) Connect(a, b *Point) {
	fmt.Printf("connecting %d,%d to %d,%d\n", a.X, a.Y, b.X, b.Y)
	if a.X == b.X {
		minY, maxY := min(a.Y, b.Y), max(a.Y, b.Y)
		for y := minY + 1; y < maxY; y++ {
			(*g)[y][a.X] = 'X'
		}
		(*g)[minY][a.X] = '#'
		(*g)[maxY][a.X] = '#'
	} else {
		minX, maxX := min(a.X, b.X), max(a.X, b.X)
		for x := minX + 1; x < maxX; x++ {
			(*g)[a.Y][x] = 'X'
		}
		(*g)[a.Y][minX] = '#'
		(*g)[a.Y][maxX] = '#'
	}
}

func (g *Grid) Fill(x, y int) {
	if x < 0 || x >= len((*g)[0]) || y < 0 || y >= len(*g) {
		return
	}
	if (*g)[y][x] != '.' {
		return
	}

	(*g)[y][x] = '#'
	g.Fill(x-1, y)
	g.Fill(x+1, y)
	g.Fill(x, y-1)
	g.Fill(x, y+1)
}

func (g *Grid) Dump() {
	for _, row := range *g {
		for _, cell := range row {
			if cell == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", cell)
			}
		}
		fmt.Println()
	}
}

type Option struct {
	Area float64
	Corner1 *Point
	Corner2 *Point
}

func NewOption(a, b *Point) *Option {
	return &Option{
		Area: (1 + math.Abs(float64(a.X - b.X))) * (1 + math.Abs(float64(a.Y - b.Y))),
		Corner1: a,
		Corner2: b,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	points := []*Point{}

	maxX := 0
	maxY := 0
	for scanner.Scan() {
		line := scanner.Text()
		point := ToPoint(strings.Split(line, ","))
		points = append(points, point)
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y > maxY {
			maxY = point.Y
		}
	}

	// maxX++
	
	data := &Grid{}
	data.Init(maxX, maxY)
	
	for idx := 0; idx < len(points) - 1; idx++ {
		data.Connect(points[idx], points[idx+ 1])
	}
	data.Connect(points[0], points[len(points)-1])
    data.Dump()

	options := []*Option{}
	for idx := 0; idx < len(points) - 1; idx++ {
		for jdx := idx + 1; jdx < len(points); jdx++ {
			options = append(options, NewOption(points[idx], points[jdx]))
		}
	}

	data.Fill(maxX-2, maxY-2)
	data.Dump()
}
