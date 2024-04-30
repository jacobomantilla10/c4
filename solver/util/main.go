package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	connectfour "github.com/jacobomantilla10/connect-four"
	"github.com/jacobomantilla10/connect-four/solver"
)

func main() {
	CreateExactOpeningsFile()
}

func CreateExactOpeningsFile() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	fmt.Println(parent)
	file, err := os.Open(filepath.Join(parent, "openings.data"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		pos, _ := strconv.Atoi(fields[0])
		mask, _ := strconv.Atoi(fields[1])
		// key, _ := strconv.Atoi(fields[2])
		val, _ := strconv.Atoi(fields[3])

		board := connectfour.MakeBoardFromOpening(uint64(pos), uint64(mask), 8)
		fmt.Printf("Solving board %d...", i)
		if val == 0 {
			lines = append(lines, fmt.Sprintf("%d %d", board.Key(), 0))
			fmt.Printf("Result is %d\n", 0)
		} else {
			res := solver.Solve(board)
			lines = append(lines, fmt.Sprintf("%d %d", board.Key(), res))
			fmt.Printf("Result is %d\n", res)
		}
		i++
	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile(filepath.Join(parent, "openingbook.data"), []byte(output), 0664)
	if err != nil {
		panic(err)
	}
}

func CreateOpeningsFile() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	file, err := os.ReadFile(filepath.Join(parent, "connect-4.data"))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		// character to the right becomes -1, 1, or 0
		fmt.Println(line)
		result := line[84:]
		val := -2
		if result == "win" {
			val = 1
		} else if result == "loss" {
			val = -1
		} else {
			val = 0
		}
		// characters to the left get turned into array and we create a board with them
		board := boardFromDataLine(strings.Split(line[:83], ","))
		// we then get the key from the board we created and that is our transposition table key
		key := board.Key()
		// then we format our file into "key value" pairs which is the value that we set our line to
		lines[i] = fmt.Sprintf("%d %d %d %d", board.Position(), board.Mask(), key, val)
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile(filepath.Join(parent, "openings.data"), []byte(output), 0664)
	if err != nil {
		panic(err)
	}
}

// Takes in a line from our connect-4.data file and converts it into a board
func boardFromDataLine(board []string) connectfour.Board {
	numMoves := 8

	var mask uint64
	var position uint64
	for i := 0; i < 7; i++ {
		for j := 0; j < 6; j++ {
			if board[(i*6)+j] == "x" {
				// update mask
				mask |= (mask + connectfour.Bottom_mask(i))
				// update board
				position |= ((1 << (i * (connectfour.HEIGHT + 1))) << j)
			} else if board[(i*6)+j] == "o" {
				// update mask
				mask |= (mask + connectfour.Bottom_mask(i))
			}
		}
	}
	return connectfour.MakeBoardFromOpening(position, mask, numMoves)
}
