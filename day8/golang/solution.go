package main

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
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

type link struct {
	distance float64
	j1_id    int
	j2_id    int
}

func calculate_shortest_links(junctures []juncture, link_number int) []link {
	perfs := [3]int{}
	lowest_links := make([]link, 0, link_number)
	var max_distance = 0.
	for i := range junctures {
		for j := i + 1; j < len(junctures); j++ {
			distance := junctures[i].distance(junctures[j])
			thislink := link{distance: distance, j1_id: i, j2_id: j}
			if len(lowest_links) < link_number {
				lowest_links = append(lowest_links, thislink)
				slices.SortFunc(lowest_links, func(a, b link) int {
					return cmp.Compare(a.distance, b.distance)
				})
				max_distance = max(max_distance, distance)
				perfs[0]++
			} else if distance < max_distance {
				lowest_links = append(lowest_links[:link_number-1], thislink)
				max_distance = max(distance, lowest_links[link_number-2].distance)
				slices.SortFunc(lowest_links, func(a, b link) int {
					return cmp.Compare(a.distance, b.distance)
				})
				perfs[1]++
			} else {
				perfs[2]++
			}
		}
	}
	fmt.Printf("Performance Shortest Links: %v\n", perfs)
	return lowest_links
}

func identify_clusters(links []link) [][]int {
	perfs := [4]int{}      // automatically zeroed
	var latest_circuit = 0 // when key is missing, Go maps return fking ZERO WTF
	circuits := make(map[int]int)
	for _, lnk := range links {
		// If you use the single return value,
		//  you get a generic "zero" if the key doesn't exists.
		j1_circuit, j1_exists := circuits[lnk.j1_id]
		j2_circuit, j2_exists := circuits[lnk.j2_id]
		if j1_exists && j2_exists {
			// This heavy branch is still the most done, around 43% of times
			// I can, if needed, use a second map for the relation val -> key
			perfs[0]++
			if j1_circuit != j2_circuit {
				// j2_circuit --> j1_circuit
				for key, val := range circuits {
					if val == j2_circuit {
						circuits[key] = j1_circuit
					}
				}
			}
		} else if j1_exists {
			perfs[1]++
			circuits[lnk.j2_id] = j1_circuit
		} else if j2_exists {
			perfs[2]++
			circuits[lnk.j1_id] = j2_circuit
		} else {
			perfs[3]++
			circuits[lnk.j1_id] = latest_circuit
			circuits[lnk.j2_id] = latest_circuit
			latest_circuit++
		}
	}
	fmt.Printf("Performance Identify Clusters Map: %v\n", perfs)
	clusters := make(map[int][]int)
	for j_id, circuit_id := range circuits {
		circuit, exists := clusters[circuit_id]
		if !exists {
			circuit = []int{j_id}
		} else {
			circuit = append(circuit, j_id)
		}
		clusters[circuit_id] = circuit
	}
	list_of_circuits := [][]int{}
	for _, val := range clusters {
		list_of_circuits = append(list_of_circuits, val)
	}
	return list_of_circuits
}

func problem1(junctures []juncture, link_number int) {
	lowest_links := calculate_shortest_links(junctures, link_number)
	circuits := identify_clusters(lowest_links)
	lengths := []int{}
	for _, circuit := range circuits {
		lengths = append(lengths, len(circuit))
	}
	// I don't really like this one because it doesn't scream "Reverse Sort" at all,
	//   even if more modern and simpler.
	slices.SortFunc(lengths, func(a, b int) int { return b - a })
	// To use the 3 greatest circuit lengths is a fixed parameter of the algorithm
	result := 1
	for i := range 3 {
		result *= lengths[i]
	}
	fmt.Printf("Problem 1: %d\n", result)
}
