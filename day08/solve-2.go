package main

import (
	"bufio"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, Z int
	Key     string
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
	z, err := strconv.Atoi(rawCoords[2])
	if err != nil {
		panic(err)
	}
	return &Point{x, y, z, fmt.Sprintf("%d,%d,%d", x, y, z)}
}

var cache = map[string]float64{}

func (p *Point) Distance(other *Point) float64 {
	p1, p2 := p, other
	if p1.Compare(p2) > 0 {
		p1, p2 = p2, p1
	}
	key := fmt.Sprintf("%s-%s", p1.Key, p2.Key)
	if val, ok := cache[key]; ok {
		return val
	}
	res := math.Sqrt(
		float64(p.X-other.X)*float64(p.X-other.X) +
			float64(p.Y-other.Y)*float64(p.Y-other.Y) +
			float64(p.Z-other.Z)*float64(p.Z-other.Z),
	)
	cache[key] = res
	return res
}

func (p *Point) Compare(other *Point) int {
	if p.X != other.X {
		return p.X - other.X
	}
	if p.Y != other.Y {
		return p.Y - other.Y
	}
	return p.Z - other.Z
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	points := []*Point{}

	for scanner.Scan() {
		line := scanner.Text()
		point := ToPoint(strings.Split(line, ","))
		points = append(points, point)
	}

	conn := make(map[string]map[string]bool)
	for _, point := range points {
		conn[point.Key] = make(map[string]bool)
	}

	// fmt.Printf("got %d points\n", len(points))

	for {

		var minDistance *float64
		var end1, end2 *Point

		for _, point := range points {
			for _, other := range points {
				if point.Key == other.Key {
					continue
				}
				distance := point.Distance(other)
				if minDistance == nil || distance < *minDistance {
					_, ok := conn[point.Key][other.Key]
					if !ok {
						minDistance = &distance
						end1 = point
						end2 = other
					}
				}
			}
		}

		// fmt.Printf("found min distance %f between %s and %s\n", *minDistance, end1.Key, end2.Key)

		if end1 == nil || end2 == nil {
			panic("end1 or end2 is nil")
		}

		conn[end1.Key][end2.Key] = true
		conn[end2.Key][end1.Key] = true

		circSeen := make(map[string]bool)

		for p := range conn {
			if circSeen[p] {
				continue
			}
			circSeen[p] = true

			seen := map[string]bool{p: true}
			queue := slices.Collect(maps.Keys(conn[p]))
			slices.Sort(queue)

			for len(queue) > 0 {
				c := queue[0]
				queue = queue[1:]
				if seen[c] {
					continue
				}
				seen[c] = true
				if circSeen[c] {
					continue
				}
				for other := range conn[c] {
					queue = append(queue, other)
				}
			}

			size := len(seen)

			if size == len(points) {
				fmt.Println(end1.X * end2.X)
				os.Exit(0)
			}
		}
	}
}
