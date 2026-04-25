package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("failed to open file with error: %s", err)
	}
	invalidSum := getInvalidSum(strings.TrimSpace(string(input)))
	log.Printf("The sum of invalid ids is %d", invalidSum)
}

func getInvalidSum(x string) int {
	idRanges := strings.Split(x, ",")
	var res int

	for _, idRange := range idRanges {
		invalidIDs := getInvalidInRange(idRange)
		for _, invalidID := range invalidIDs {
			val, err := strconv.Atoi(invalidID)
			if err != nil {
				log.Fatalf("Invalid input '%s' for invalid id '%s'", invalidID, x)
			}
			res += val
		}
	}

	return res
}

func getInvalidInRange(x string) []string {
	val := strings.Split(x, "-")
	if len(val) != 2 {
		log.Fatalf("Invalid input '%s'", x)
	}
	fromStr := val[0]
	from, err := strconv.Atoi(fromStr)
	if err != nil {
		log.Fatalf("Invalid input '%s' for from part of '%s'", fromStr, x)
	}
	toStr := val[1]
	to, err := strconv.Atoi(val[1])
	if err != nil {
		log.Fatalf("Invalid input '%s' for to part of '%s'", toStr, x)
	}
	result := make([]string, 0)
	for i := from; i <= to; i++ {
		input := strconv.Itoa(i)
		if !valid(input) {
			result = append(result, input)
		}
	}
	return result
}

func valid(x string) bool {
	xLen := len(x)

	if xLen == 1 || math.Mod(float64(xLen), 2) == 1 {
		return true
	} else if x[0:1] == "0" {
		return false
	}

	// for i := range x {
	// 	for scope := 1; scope <= i; scope++ {
	// 		if i-scope < 0 || i+scope > xLen {
	// 			continue
	// 		}
	//
	// 		before := x[i-scope : i]
	// 		after := x[i : i+scope]
	// 		if before == after {
	// 			return false
	// 		}
	// 	}
	// }
	//
	// return true

	return x[0:xLen/2] != x[xLen/2:xLen]
}
