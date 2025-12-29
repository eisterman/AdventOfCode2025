package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func main() {
	path := filepath.Join("/home/fpasqua/git/AdventOfCode2025/day4/input.data")
	// path := filepath.Join("/home/fpasqua/git/AdventOfCode2025/day4/test.data")
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(dat), "\n") // table[y][x]
	// Remove trailing empty row
	var table = content[:len(content)-1]
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
				rowRunes := []rune(table[y])
				rowRunes[x] = 'x'
				table[y] = string(rowRunes)
			}
		}
		// fmt.Print("\n")
	}
	fmt.Printf("Problema 1: %d", res)
	// Problem 2 - remove the removable and repeat
}
