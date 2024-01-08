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

func game_is_valid(game []string) bool {
    for _, result := range game {
        dice := strings.Split(result, ",")

        for _, die := range dice {
            num, color := delete_empty(strings.Split(die, " "))[0], delete_empty(strings.Split(die, " "))[1]
            
            numint, err := strconv.Atoi(num)
            if err != nil {
                fmt.Println("Error converting string to int")
                return false
            }

            if color == "blue" && numint > 14 {
                return false
            } else if color == "red" && numint > 12 {
                return false
            } else if color == "green" && numint > 13 {
                return false
            }
        }
    }

    return true
}

func get_power(game []string) int {
    var min_green int;
    var min_blue int;
    var min_red int;

    for _, result := range game {
        dice := strings.Split(result, ",")

        for _, die := range dice {
            num, color := delete_empty(strings.Split(die, " "))[0], delete_empty(strings.Split(die, " "))[1]

            numint, err := strconv.Atoi(num)
            if err != nil {
                fmt.Println("Error converting string to int")
            }
            
            if color == "green" && numint > min_green {
                min_green = numint
            }
            if color == "blue" && numint > min_blue {
                min_blue = numint
            }
            if color == "red" && numint > min_red {
                min_red = numint
            }
        }
    }

    return min_green * min_blue * min_red
}


func main() {
    contents, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Println("Error reading file")
    }

    valid_game_ids := make([]string, 0)
    powers := make([]int, 0)

    games := delete_empty(strings.Split(string(contents), "\n"))

    for _, game := range games {
        split_game := strings.Split(game, ":")
        id := strings.Split(split_game[0], " ")[1]

        results := delete_empty(strings.Split(split_game[1], ";"))
        is_valid := game_is_valid(results)

        power := get_power(results)
        powers = append(powers, power)

        if is_valid {
            valid_game_ids = append(valid_game_ids, id)
        }
    } 

    fmt.Println("Valid game ids: ", valid_game_ids)

    id_sum := 0

    for _, id := range valid_game_ids {
        idint, err := strconv.Atoi(id)
        if err != nil {
            fmt.Println("Error converting string to int")
        }

        id_sum += idint
    }

    fmt.Println("Sum of valid game ids: ", id_sum)
    fmt.Println("")

    fmt.Println("Powers: ", powers)

    power_sum := 0

    for _, power := range powers {
        power_sum += power
    }

    fmt.Println("Sum of powers: ", power_sum)
}
