package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/albertorestifo/dijkstra"
)

func numbers(raw string) []int {
	parts := strings.Split(raw, ",")
	numbers := make([]int, len(parts))
	for i, part := range parts {
		numbers[i], _ = strconv.Atoi(part)
	}
	// fmt.Printf("numbers(%s) = %v\n", raw, numbers)
	return numbers
}

func key(v []int) string {
	str := []string{}
	for _, v := range v {
		str = append(str, strconv.Itoa(v))
	}
	return strings.Join(str, ",")
}

func exceeds(terminus, curr string) bool {
	t := numbers(terminus)
	c := numbers(curr)
	for i := range t {
		if c[i] < t[i] {
			return false
		}
	}
	return true
}

func gap(terminus, curr string, button []int) bool {
	t := numbers(terminus)
	c := numbers(curr)
	for _, pos := range button {
		if c[pos] < t[pos] {
			return true
		}
	}
	return false
}

func apply(curr string, button []int) string {
	c := numbers(curr)
	for _, pos := range button {
		c[pos]++
	}
	res := key(c)
	// fmt.Printf("curr = %s, button = %v, next = %s\n", curr, button, res)
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lights string
	var joltage []int
	var buttons [][]int

	sum := 0

	for scanner.Scan() {
		raw := scanner.Text()
		pieces := strings.Split(raw, " ")
		lights = strings.Trim(pieces[0], "[]")
		pieces = pieces[1:]
		joltage = numbers(strings.Trim(pieces[len(pieces)-1], "{}"))
		pieces = pieces[:len(pieces)-1]
		for _, button := range pieces {
			buttons = append(buttons, numbers(strings.Trim(button, "()")))
		}

		origin := strings.Join(strings.Split(strings.Repeat("0", len(lights)), ""), ",")
		terminus := key(joltage)
		fmt.Printf("joltage = %v, buttons = %+v\n", joltage, buttons)
		fmt.Printf("origin = %s, terminus = %s\n", origin, terminus)

		queue := []string{origin}
		g := map[string]map[string]int{}
		seen := map[string]bool{}

		// fmt.Printf("Filling out graph.\n")

	OUTER:
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			if seen[curr] {
				continue
			}
			seen[curr] = true
			if curr == terminus {
				break
			}
			if exceeds(terminus, curr) {
				continue
			}
			for _, button := range buttons {
				c := curr
				cost := 1
				for gap(terminus, c, button) {
					next := apply(c, button)
					if _, ok := g[c]; !ok {
						g[c] = map[string]int{}
					}
					g[c][next] = 1
					if (len(g)) % 1000 == 0 {
						fmt.Printf("graph size = %d\n", len(g))
					}
					if _, ok := g[curr][next]; !ok || g[curr][next] > cost {
						// g[curr][next] = cost
					}
					cost++
					c = next
					if !seen[next] {
						queue = append(queue, next)
					}
					if next == terminus {
						break OUTER
					}
				}
			}
		}

		if _, ok := g[terminus]; !ok {
			g[terminus] = map[string]int{}
		}
		g[terminus][terminus] = 1

		// dump, _ := json.Marshal(g)
		// fmt.Println(string(dump))

		var graph dijkstra.Graph = g

		// dump, _ = json.Marshal(g)
		// fmt.Println(string(dump))

		fmt.Printf("Traversing graph of size %d.\n", len(g))
		path, cost, err := graph.Path(origin, terminus)
		if err != nil {
			panic(err)
		}

		fmt.Printf("path = %+v, cost = %d\n", path, cost)

		sum += cost
	}

	fmt.Printf("\nsum = %d\n\n", sum)
}
