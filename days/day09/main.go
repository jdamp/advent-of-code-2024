package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input string

const FREEBLOCK = "."

// Parses the input string into the individual blocks
func parseInput(input string) []string {
	var result []string
	var id string
	for i, char := range input {
		if i%2 == 0 {
			id = fmt.Sprintf("%d", i/2)
		} else {
			id = FREEBLOCK
		}
		num := int(char - 0x0030)
		for j := 0; j < num; j++ {
			result = append(result, id)
		}
	}
	return result
}

func moveBlocks(blocks []string) []string {
	first := 0
	last := len(blocks) - 1
	for first < last {
		// If we are not on a free block move further
		if blocks[first] != FREEBLOCK {
			first += 1
			continue
		}
		if blocks[last] == FREEBLOCK {
			last -= 1
			continue
		}
		// Move the block from last to first
		blocks[first] = blocks[last]
		blocks[last] = FREEBLOCK
	}
	return blocks
}

func moveFiles(blocks []string) []string {
	// We now need four pointers instead of two
	startFirst, endFirst := 0, 0
	startLast, endLast := len(blocks)-1, len(blocks)-1
	// Identify the first file block
	// Next, move the last pointers to reach the last file

	for startLast > 0 {
		startFirst, endFirst = 0, 0
		for blocks[endFirst] == FREEBLOCK {
			endLast--
		}
		fileId := blocks[endLast]
		endLast = startLast
		for startLast > 0 && blocks[startLast] == fileId {
			startLast--
		}
		for startLast > endFirst {
			// First, move the first pointers to reach the first free block
			for blocks[startFirst] != FREEBLOCK {
				startFirst++
				endFirst = startFirst
			}
			for blocks[endFirst] == FREEBLOCK && endFirst < endLast {
				endFirst++
			}
			// Now we can check whether the file fits into the free blocks
			if endFirst-startFirst >= endLast-startLast {
				for delta := 0; delta < endLast-startLast; delta++ {

					blocks[startFirst+delta] = blocks[startLast+delta+1]
					blocks[startLast+delta+1] = FREEBLOCK
				}

				startFirst, endFirst = 0, 0

				break // Move to the next file
			} else {
				// File doesn't fit move block pointers
				startFirst = endFirst
			}
		}
	}
	return blocks
}

func part1(input string) (checksum int) {
	disk := parseInput(input)
	blocks := moveBlocks(disk)
	for pos, id := range blocks {
		val, _ := strconv.Atoi(id)
		checksum += val * pos
	}
	return checksum
}

func part2(input string) (checksum int) {
	disk := parseInput(input)
	blocks := moveFiles(disk)
	for pos, id := range blocks {
		val, _ := strconv.Atoi(id)
		checksum += val * pos
	}
	return checksum
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "run which part (1 or 2)")
	flag.Parse()
	if part == 1 {
		result := part1(input)
		fmt.Println("Result: ", result)
	} else {
		result := part2(input)
		fmt.Println("Result: ", result)
	}

}
