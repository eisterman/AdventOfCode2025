package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Unused but good example of generics
func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func PrintTable(table [][]rune) {
	for _, row := range table {
		for _, val := range row {
			fmt.Printf("%c", val)
		}
		fmt.Print("\n")
	}
}

func PrintNeigh(table [][]rune) {
	for _, row := range table {
		for _, val := range row {
			fmt.Printf("%d", val)
		}
		fmt.Print("\n")
	}
}

func prob1(table [][]rune) int {
	maxY := len(table)
	maxX := len(table[0])
	neigh := make([][]int, maxY)
	for i := range neigh {
		neigh[i] = make([]int, maxX)
		// all zeros
	}
	for y, row := range table {
		for x, val := range row {
			// fmt.Printf("(%d,%d)%c ", x, y, val)
			if val == '@' {
				for _, dx := range []int{-1, 0, 1} {
					for _, dy := range []int{-1, 0, 1} {
						if dx == 0 && dy == 0 {
							continue
						}
						accX := x + dx
						accY := y + dy
						if accX < 0 || accX >= maxX || accY < 0 || accY >= maxY {
							continue
						}
						neigh[accY][accX] += 1
					}
				}
			}
		}
		// fmt.Print("\n")
	}
	var res = 0
	for y, row := range neigh {
		for x, val := range row {
			// fmt.Printf("(%d,%d)%c ", x, y, val)
			if table[y][x] == '@' {
				if val < 4 {
					res += 1
				}
				// table[y][x] = 'x'
			}
		}
		// fmt.Print("\n")
	}
	return res
}

func prob2(table [][]rune) int {
	maxY := len(table)
	maxX := len(table[0])
	// Initialize neighbours
	neigh := make([][]int, maxY)
	for i := range neigh {
		neigh[i] = make([]int, maxX)
		// all zeros
	}
	// PrintTable(table)
	for y, row := range table {
		for x, val := range row {
			if val == '@' {
				for _, dx := range []int{-1, 0, 1} {
					for _, dy := range []int{-1, 0, 1} {
						if dx == 0 && dy == 0 {
							continue
						}
						accX := x + dx
						accY := y + dy
						if accX < 0 || accX >= maxX || accY < 0 || accY >= maxY {
							continue
						}
						neigh[accY][accX] += 1
					}
				}
			}
		}
	}
	// Iterate removal
	var res = 0
	for {
		var dres = 0
		// Find removable
		for y, row := range neigh {
			for x, val := range row {
				// fmt.Printf("(%d,%d)%c ", x, y, val)
				if table[y][x] == '@' {
					if val < 4 {
						dres += 1
						table[y][x] = 'x'
					}
				}
			}
			// fmt.Print("\n")
		}
		// Exit if nothing changed
		if dres == 0 {
			break
		}
		res += dres
		// Actually remove them
		for y, row := range table {
			for x, val := range row {
				// fmt.Printf("(%d,%d)%c ", x, y, val)
				if val == 'x' {
					for _, dx := range []int{-1, 0, 1} {
						for _, dy := range []int{-1, 0, 1} {
							if dx == 0 && dy == 0 {
								continue
							}
							accX := x + dx
							accY := y + dy
							if accX < 0 || accX >= maxX || accY < 0 || accY >= maxY {
								continue
							}
							neigh[accY][accX] -= 1
						}
					}
					table[y][x] = '.'
				}
			}
			// fmt.Print("\n")
		}
	}
	return res
}

func main() {
	path := filepath.Join("/home/fpasqua/git/AdventOfCode2025/day4/input.data") // 1540 - 8972
	// path := filepath.Join("/home/fpasqua/git/AdventOfCode2025/day4/test.data") // 13 - 43
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(dat), "\n") // table[y][x]
	// Remove trailing empty row
	var table = make([][]rune, len(content)-1)
	for y, row := range content[:len(content)-1] {
		table[y] = []rune(row)
	}
	// var table = content[:len(content)-1]
	res1 := prob1(table)
	fmt.Printf("Problema 1: %d\n", res1)
	// Problem 2 - remove the removable and repeat
	res2 := prob2(table)
	fmt.Printf("Problema 2: %d\n", res2)
}
