package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
    time, distance := parseFile(file)

    numberOfWins := 0
    for j := 1; j <= time; j++ {
        distanceTraveled := j * (time - j)
        if distanceTraveled > distance {
            numberOfWins++
        }
    }

    fmt.Println(numberOfWins)
}

func parseFile(file []byte) (int, int) {
    var time int
    var timeString string
    var distance int 
    var distanceString string
    lines := delete_empty(strings.Split(string(file), "\n"))

    timeStrings := delete_empty(strings.Split(strings.Split(lines[0], ":")[1], " "))
    for _, s := range timeStrings {
        timeString += s
    }
    time, _ = strconv.Atoi(timeString)

    distanceStrings := delete_empty(strings.Split(strings.Split(lines[1], ":")[1], " "))
    for _, s := range distanceStrings {
        distanceString += s
    }
    distance, _ = strconv.Atoi(distanceString)

    return time, distance
}
