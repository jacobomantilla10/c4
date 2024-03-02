package solver

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	connectfour "github.com/jacobomantilla10/connect-four"
)

func TestMiniMax(t *testing.T) {
	type test struct {
		name  string
		input connectfour.Board
		want  int
	}
	var tests []test

	file, err := os.Open("testfiles/Test_L3_R1")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		boardString := fields[0]
		want, _ := strconv.Atoi(fields[1])
		board, _ := connectfour.MakeBoardFromString(boardString)
		tests = append(tests, test{fmt.Sprintf("%s should be %d", boardString, want), board, want})
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := MiniMax(tt.input, 14, -1000, 1000, 1)
			if ans != tt.want {
				t.Errorf("at item %d got %d want %d", i+1, ans, tt.want)
			}
		})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
