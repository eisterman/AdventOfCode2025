package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type juncture struct {
	x int
	y int
	z int
}

func (j juncture) sub(rhs juncture) juncture {
	return juncture{x: j.x - rhs.x, y: j.y - rhs.y, z: j.z - rhs.z}
}

func (j juncture) length() float64 {
	return math.Sqrt(float64(j.x*j.x + j.y*j.y + j.z*j.z))
}

func (j juncture) distance(rhs juncture) float64 {
	return j.sub(rhs).length()
}

func ParseRow(s string) (juncture, error) {
	var result juncture
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		return result, errors.New("Can manage only 3 numbers")
	}
	a, e := strconv.Atoi(parts[0])
	if e != nil {
		return result, e
	}
	b, e := strconv.Atoi(parts[1])
	if e != nil {
		return result, e
	}
	c, e := strconv.Atoi(parts[2])
	if e != nil {
		return result, e
	}
	result.x = a
	result.y = b
	result.z = c
	return result, nil
}

func main() {
	path := filepath.Join("/home/fpasqua/git/AdventOfCode2025/day8/input.data")
	link_number := 1000
	// path := filepath.Join("/home/fpasqua/git/AdventOfCode2025/day8/test.data")
	// link_number := 10
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := strings.Split(strings.TrimSpace(string(dat)), "\n") // table[y][x]
	var table = make([]juncture, len(content))
	for y, row := range content {
		x, _ := ParseRow(row)
		table[y] = x
	}
	problem1(table, link_number)
}

type jlink struct {
	j1_id int
	j2_id int
}

type circuit struct {
	juncture_ids []int
}

func extract_related(junctures []juncture, links []jlink) []circuit {
	var to_distribute []int
	for i := range junctures {
		to_distribute = append(to_distribute, i)
	}
	var result []circuit
	for len(to_distribute) > 0 {
		junc_id := to_distribute[0]
		var associated = []int{junc_id}
		var changed = true
		for changed {
			changed = false
			for _, link := range links {
				j1_contained := slices.Contains(associated, link.j1_id)
				j2_contained := slices.Contains(associated, link.j2_id)
				if j1_contained && j2_contained {
					// If both are already associated, nothing to do
					continue
				} else if j1_contained {
					associated = append(associated, link.j2_id)
					changed = true
				} else if j2_contained {
					associated = append(associated, link.j1_id)
					changed = true
				}
			}
		}
		// Create circuit
		result = append(result, circuit{juncture_ids: associated})
		// Remove distributed from to_distribute
		var c []int
		for _, idx := range to_distribute {
			if !slices.Contains(associated, idx) {
				c = append(c, idx)
			}
		}
		to_distribute = c
	}
	return result
}

func problem1(junctures []juncture, link_number int) {
	linkSet := make(map[jlink]bool)
	for range link_number {
		var j1_id, j2_id int
		var distance = math.MaxFloat64
		for i, jun1 := range junctures {
			for j := i + 1; j < len(junctures); j++ {
				jun2 := junctures[j]
				if linkSet[jlink{j1_id: i, j2_id: j}] {
					continue
				}
				d := jun1.distance(jun2)
				if d < distance {
					distance = d
					j1_id = i
					j2_id = j
				}
			}
		}
		linkSet[jlink{j1_id, j2_id}] = true
	}
	links := make([]jlink, 0, len(linkSet)) // type, length, capacity
	for link := range linkSet {
		links = append(links, link)
	}
	// Identify clusters
	fmt.Printf("Links: %v\n", links)
	var circuits = extract_related(junctures, links)
	for i, c := range circuits {
		fmt.Printf("Circuit %d has %d junctions %v\n", i, len(c.juncture_ids), c.juncture_ids)
	}
	// Result is size of 3 greatest circuits
	var sizes []int
	for _, c := range circuits {
		sizes = append(sizes, len(c.juncture_ids))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	var result = 1
	for i, length := range sizes {
		if i >= 3 {
			break
		}
		result *= length
	}
	fmt.Printf("Problem 1: %d\n", result)
}
