package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func delete_empty(s []string) []string {
    var r []string
    for _, str := range s {
        if str != "" {
            r = append(r, str)
        }
    }
    return r
}

func main() {
    file, _ := os.ReadFile("input.txt")
    tree := strings.Split(string(file), "\n\n")
    seed_ranges := get_seeds(strings.Split(tree[0], ":")[1])
    mappings := create_range_map(tree)

    var wg sync.WaitGroup
    
    var location int
    for _, seed_range := range seed_ranges {
        wg.Add(1)
        go func(seed_range []int) {
            defer wg.Done()
            fmt.Println("Checking range: ", seed_range)
            max_seed := seed_range[0] + (seed_range[1] -1)
            for seed := seed_range[0]; seed <= max_seed; seed+=1 {
                get_location(1, seed, &mappings, &location)
            }
        }(seed_range)
    }

    wg.Wait()

    fmt.Println(location)
}

func get_seeds(seed_string string) [][]int {
    seeds := delete_empty(strings.Split(seed_string, " "))
    seed_ranges := [][]int{}

    next_range := []int{}
    for i, seed := range seeds {
        num, _ := strconv.Atoi(seed)
        next_range = append(next_range, num)

        if i != 0 && i%2 != 0 {
            seed_ranges = append(seed_ranges, next_range)
            next_range = []int{}
        }
    }

    return seed_ranges
}

func get_location(index int, value int, mapping *map[int][]RangeMap, location *int) {
    new_val := value
    for _, range_map := range (*mapping)[index] {
        if range_map.Source <= value && value <= range_map.Source + (range_map.Range - 1) {
            diff := value - range_map.Source
            new_val = range_map.Destination + diff 
            break
        }
    }
    if len((*mapping)[index + 1]) != 0 {
        get_location(index + 1, new_val, mapping, location)
    } else {
        if *location == 0 || new_val < *location {
            *location = new_val
        }
    }
}

func create_range_map(tree []string) map[int][]RangeMap {
    new_map := map[int][]RangeMap{}
    for i, row := range tree {
        if i == 0 {
            continue
        }

        var nums []RangeMap
        for _, line := range strings.Split(row, "\n") {
            numstrings := strings.Split(line, " ")
            numstrings = delete_empty(numstrings)
            if len(numstrings) != 3 {
                continue
            }

            var newmap RangeMap

            for i, num := range numstrings {
                numint, _ := strconv.Atoi(num)
                switch i {
                case 0:
                    newmap.Destination = numint
                case 1:
                    newmap.Source = numint
                default:
                    newmap.Range = numint
                }
            }

            nums = append(nums, newmap)
        }

        new_map[i] = nums
    }
    return new_map
}

type RangeMap struct {
    Destination int
    Source      int
    Range       int
}
