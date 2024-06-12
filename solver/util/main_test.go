package main

import (
	"fmt"
	"testing"

	"github.com/jacobomantilla10/connect-four/game"
)

func TestBoardFromDataLine(t *testing.T) {
	b1, _ := game.MakeBoardFromString("33444444")
	b2, _ := game.MakeBoardFromString("35444444")
	b3, _ := game.MakeBoardFromString("32444444")
	var tests = []struct {
		input []string
		want  uint64
	}{
		{[]string{"b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "x", "o", "b", "b", "b", "b", "x", "o", "x", "o", "x", "o", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"}, b1.Key()},
		{[]string{"b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "x", "b", "b", "b", "b", "b", "x", "o", "x", "o", "x", "o", "o", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"}, b2.Key()},
		{[]string{"b", "b", "b", "b", "b", "b", "o", "b", "b", "b", "b", "b", "x", "b", "b", "b", "b", "b", "x", "o", "x", "o", "x", "o", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"}, b3.Key()},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			b := boardFromDataLine(tt.input)
			res := b.Key()
			if res != tt.want {
				t.Errorf("At item %d got %b want %b", i, res, tt.want)
			}
		})
	}
}
