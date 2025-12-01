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

		rotations := int(dist / 100)
		remainder := dist - 100*rotations
		full_rotation := remainder == 0
		at_zero := curr == 0

		password += rotations

		if at_zero {
			if !full_rotation {
				if dir == 'L' {
					curr = 100 - remainder
				} else {
					curr = remainder
				}
			}
		} else if !full_rotation {
			if dir == 'L' {
				curr -= remainder
			} else {
				curr += remainder
			}
			wrapped := false
			if curr < 0 {
				curr += 100
				password++
				wrapped = true
			}
			if curr >= 100 {
				curr -= 100
				password++
				wrapped = true
			}
			if curr == 0 && !wrapped {
				password++
			}
		}

		fmt.Printf("The dial is rotated %s to point at %d, and the current password is %d.\n", line, curr, password)
	}
}
