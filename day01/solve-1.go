package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	curr := 50
	password := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if dir == 'L' {
			curr -= dist
			for curr < 0 {
				curr += 100
			}
		} else {
			curr += dist
			for curr >= 100 {
				curr -= 100
			}
		}

		if curr == 0 {
			password++
		}

		fmt.Printf("The dial is rotated %s to point at %d, and the current password is %d.\n", line, curr, password)
	}
}
