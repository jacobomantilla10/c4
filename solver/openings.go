package solver

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// create openingBook which is the same type as transposition
type OpeningBook struct {
	Openings map[uint64]int8
	Size     int
}

// create function to initialize it using the data file
func MakeOpeningBook() OpeningBook {
	openings := map[uint64]int8{}
	size := 0

	file, err := os.Open("openingbook.data")
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

		openings[uint64(key)] = int8(val)
		size++
	}

	return OpeningBook{openings, size}
}

// Get opening if it exists, return opening with -30 value if it does not.
func (o OpeningBook) Opening(key uint64) int8 {
	val, ok := o.Openings[key]
	if !ok {
		// return opening with -30 val to signify it is not valid
		return -30
	}
	return val
}
