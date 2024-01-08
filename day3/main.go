package main

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
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

func is_number(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func is_period(s string) bool {
	return s == "."
}

func is_asterisk(s string) bool {
    return s == "*"
}

func check_surrounding_chars(i int, j int, matrix [][]string) bool {
	surrounding_chars := make([]string, 0)

	if i > 0 {
		surrounding_chars = append(surrounding_chars, matrix[i-1][j])
		if j > 0 {
			surrounding_chars = append(surrounding_chars, matrix[i-1][j-1])
		}
		if j+1 < len(matrix[i]) {
			surrounding_chars = append(surrounding_chars, matrix[i-1][j+1])
		}
	}

	if j > 0 {
		surrounding_chars = append(surrounding_chars, matrix[i][j-1])
	}

	if j+1 < len(matrix) {
		surrounding_chars = append(surrounding_chars, matrix[i][j+1])
	}

	if i+1 < len(matrix) {
		surrounding_chars = append(surrounding_chars, matrix[i+1][j])
		if j > 0 {
			surrounding_chars = append(surrounding_chars, matrix[i+1][j-1])
		}
		if j+1 < len(matrix[i]) {
			surrounding_chars = append(surrounding_chars, matrix[i+1][j+1])
		}
	}

	for _, char := range surrounding_chars {
		if !is_number(char) && !is_period(char) {
            return true
        }
	}

	return false
}

func main() {
	contents, _ := os.ReadFile("input.txt")

	lines := delete_empty(strings.Split(string(contents), "\n"))

	matrix := make([][]string, 0)

	for _, line := range lines {
		line_array := delete_empty(strings.Split(line, ""))
		matrix = append(matrix, line_array)
	}

	var numbers []float64

	for i, line := range matrix {
		should_build := true

		for j, char := range line {
			var built_num string
			if is_number(char) {
				if should_build {
					built_num += char
					next_num := j + 1

					is_part_number := check_surrounding_chars(i, j, matrix)

					for should_build {
						if next_num < len(line) && is_number(line[next_num]) {
							built_num += line[next_num]
							if !is_part_number {
								is_part_number = check_surrounding_chars(i, next_num, matrix)
							}
							next_num += 1
						} else {
							should_build = false
						}
					}

					num, _ := strconv.ParseFloat(built_num, 64)

					if is_part_number {
						numbers = append(numbers, num)
					}

					built_num = ""
				}
			} else {
				should_build = true
			}
		}
	}

    var sum float64 = 0

    for _, num := range numbers {
        sum += num
    }

    fmt.Println(sum)
    
    grid := map[image.Point]rune{}
    for y, s := range strings.Fields(string(contents)) {
        for x, r := range s {
            if r != '.' && !unicode.IsDigit(r) {
                grid[image.Point{x, y}] = r
            }
        }
    }

    parts := map[image.Point][]int{}
    for y, s := range strings.Fields(string(contents)) {
        for _, m := range regexp.MustCompile(`\d+`).FindAllStringIndex(s, -1) {
            bounds := map[image.Point]struct{}{}
            for x := m[0]; x < m[1]; x++ {
                for _, d := range []image.Point{
                    {-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
                } {
                    bounds[image.Point{x, y}.Add(d)] = struct{}{}
                }
            }

            n, _ := strconv.Atoi(s[m[0]:m[1]])
            for p := range bounds {
                if _, ok := grid[p]; ok {
                    parts[p] = append(parts[p], n)
                }
            }
        }
    }

    part2 := 0
    for p, ns := range parts {
        if grid[p] == '*' && len(ns) == 2 {
            part2 += ns[0] * ns[1]
        }
    }

    fmt.Println(part2)
}
