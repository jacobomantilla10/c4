package solver

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Opening struct {
	val  int8
	flag uint8
}

// create openingBook which is the same type as transposition
type OpeningBook struct {
	Openings map[uint64]Opening
	Size     int
}

// create function to initialize it using the data file
func MakeOpeningBook() OpeningBook {
	openings := map[uint64]Opening{}
	size := 0

	file, err := os.Open("openings.data")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		key, _ := strconv.Atoi(fields[0])
		val, _ := strconv.Atoi(fields[1])
		var flag uint8
		if val == -1 {
			flag = UPPER
		} else if val == 1 {
			flag = LOWER
		} else {
			flag = EXACT
		}
		openings[uint64(key)] = Opening{int8(val), flag}
		size++
	}

	return OpeningBook{openings, size}
}

// Get opening if it exists, return opening with -2 value if it does not.
func (o OpeningBook) Opening(key uint64) Opening {
	opening, ok := o.Openings[key]
	if !ok {
		// return opening with -2 val to signify it is not valid
		return Opening{-2, 0}
	}
	return Opening{opening.val, opening.flag}
}
