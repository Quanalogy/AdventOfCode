package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("failed to open file with error: %s", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var dials []Dial
	for scanner.Scan() {
		line := scanner.Text()
		dials = append(dials, NewDial(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while reading file: %s", err)
	}

	if len(dials) != 4099 {
		log.Fatalf("expected 4099 dials, got %d", len(dials))
	}

	log.Printf("The dial password part1 is: %d", DialPassword(dials, 50, true))
	log.Printf("The dial password part2 is: %d", DialPassword(dials, 50, false))
}

type Direction int

const (
	DialMax int       = 100
	Left    Direction = iota
	Right
)

type Dial struct {
	Direction Direction
	Amount    int
}

func NewDial(input string) Dial {
	amount, err := strconv.Atoi(input[1:])
	if err != nil {
		log.Fatalf("failed to convert amount with error: %s", err)
	}

	if input[0] == 'L' {
		return Dial{Direction: Left, Amount: amount}
	}
	return Dial{Direction: Right, Amount: amount}
}

func (d *Dial) Dial(current int) (res int) {
	// In this method I prefer a solution where I get the proper dial value
	// over solutions where we can go between -99 and 99 because it makes more sense to me.
	// It does however impact the solution here as we need to do crazy things below to make it wrap around properly.
	if d.Direction == Left {
		res = (current - d.Amount) % DialMax
		if res < 0 {
			return res + DialMax
		}
	} else {
		res = current + d.Amount
	}
	return int(float64(res % DialMax))
}

func (d *Dial) DialZeroCrossings(current int) (crossings int) {
	crossings = int(d.Amount / DialMax)
	remainder := d.Amount % DialMax
	if current == 0 {
		return crossings
	}
	if d.Direction == Left {
		if d.Amount < current {
			return 0
		}
		if current-remainder <= 0 {
			crossings += 1
		}
	} else {
		if d.Amount+current < DialMax {
			return 0
		}
		if current+remainder >= DialMax {
			crossings += 1
		}
	}

	return crossings
}

func DialPassword(dials []Dial, current int, part1 bool) int {
	currentDialValue := current
	password := 0
	for _, dial := range dials {
		if !part1 {
			// needs to be done before dialing
			password += dial.DialZeroCrossings(currentDialValue)
		}
		currentDialValue = dial.Dial(currentDialValue)
		if part1 {
			if currentDialValue == 0 {
				password += 1
			}
		}
	}
	return password
}
