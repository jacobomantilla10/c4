package solver

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jacobomantilla10/connect-four/game"
)

type test struct {
	name  string
	input game.Board
	want  int
}

func createTable(fn string) []test {
	var tests []test

	file, err := os.Open(fn)
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
		board, _ := game.MakeBoardFromString(boardString)
		tests = append(tests, test{fmt.Sprintf("%s should be %d", boardString, want), board, want})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return tests
}

func TestEasyEasy(t *testing.T) {
	tests := createTable("testfiles/Test_L3_R1")
	var totalTime time.Duration
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			ans := Solve(tt.input)
			end := time.Now()
			totalTime += end.Sub(start)
			if ans != tt.want {
				t.Errorf("at item %d got %d want %d", i+1, ans, tt.want)
			}
		})
	}
	t.Logf("\nTotal time: %s, Mean time: %s", totalTime.Round(time.Microsecond), (totalTime / 1000).Round(time.Microsecond))
}

func TestMediumEasy(t *testing.T) {
	tests := createTable("testfiles/Test_L2_R1")
	var totalTime time.Duration
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			ans := Solve(tt.input)
			// _, ans := GetBestMove(tt.input)
			end := time.Now()
			totalTime += end.Sub(start)
			if ans != tt.want {
				t.Errorf("at item %d got %d want %d", i+1, ans, tt.want)
			}
		})
	}
	t.Logf("\nTotal time: %s, Mean time: %s", totalTime.Round(time.Millisecond), (totalTime / 1000).Round(time.Millisecond))
}

func TestMediumMedium(t *testing.T) {
	tests := createTable("testfiles/Test_L2_R2")
	var totalTime time.Duration
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			ans := Solve(tt.input)
			end := time.Now()
			totalTime += end.Sub(start)
			if ans != tt.want {
				t.Errorf("at item %d got %d want %d", i+1, ans, tt.want)
			}
		})
	}
	t.Logf("\nTotal time: %s, Mean time: %s", totalTime.Round(time.Second), (totalTime / 1000).Round(time.Second))
}

func TestHardEasy(t *testing.T) {
	tests := createTable("testfiles/Test_L1_R1")
	var totalTime time.Duration
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			ans := Solve(tt.input)
			end := time.Now()
			totalTime += end.Sub(start)
			if ans != tt.want {
				t.Errorf("at item %d got %d want %d", i+1, ans, tt.want)
			}
		})
	}
	t.Logf("\nTotal time: %s, Mean time: %s", totalTime.Round(time.Second), (totalTime / 1000).Round(time.Second))
}

func TestHardMedium(t *testing.T) {
	tests := createTable("testfiles/Test_L1_R2")
	var totalTime time.Duration
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			ans := Solve(tt.input)
			end := time.Now()
			totalTime += end.Sub(start)
			if ans != tt.want {
				t.Errorf("at item %d got %d want %d", i+1, ans, tt.want)
			}
		})
	}
	t.Logf("\nTotal time: %s, Mean time: %s", totalTime.Round(time.Second), (totalTime / 1000).Round(time.Second))
}

func TestHardHard(t *testing.T) {
	tests := createTable("testfiles/Test_L1_R3")
	var totalTime time.Duration
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			ans := Solve(tt.input)
			// _, ans := GetBestMove(tt.input)
			end := time.Now()
			totalTime += end.Sub(start)
			if ans != tt.want {
				t.Errorf("at item %d got %d want %d", i+1, ans, tt.want)
			}
		})
	}
	t.Logf("\nTotal time: %s, Mean time: %s", totalTime.Round(time.Second), (totalTime / 1000).Round(time.Second))
}
