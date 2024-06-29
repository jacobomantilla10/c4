package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"

	"github.com/jacobomantilla10/connect-four/game"
)

func main() {
	// call read book to get book saved
	book := Read_book("bookDeepDist.dat")
	fmt.Println(book[0])
	slices.SortFunc(book, func(a, b []int) int {
		return a[0] - b[0]
	})
	// use To_huffman from notebook and To_huffman from board to check if they are the same
	board, _ := game.MakeBoardFromString("444442222666")
	// use get book value to find the value of a position using the book
	fmt.Println(Get_book_value(board, book)) // should print 71
}

func Read_book(filename string) [][]int {
	// initialize list
	book := [][]int{}
	buffer1Size := 4 // Used for first 3 bytes which give us the position
	buffer2Size := 1 // Used for last byte which gives us the score
	// open file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer1 := make([]byte, buffer1Size)
	buffer2 := make([]byte, buffer2Size)

	// read line and append to list until there are no more entries
	for {
		_, err := file.Read(buffer1)

		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			break // reached EOF
		}
		// create integer from bytes using big byte order and signed number. This is our huffman encoded position
		//pos := binary.BigEndian.Uint32(buffer1)
		pos := int(binary.BigEndian.Uint32(buffer1))
		// read one more byte into buffer2
		file.Read(buffer2)
		// create integer from bytes using big byte order and signed number. This is our score
		score1 := int8(buffer2[0]) // not sure if need to convert here
		// append this position to the book
		var score int
		if score1 < 0 {
			distance_turns := (int(score1) + 100 + 1) / 2
			score = (22 - (6 + distance_turns)) * -1
		} else {
			distance_turns := (100 - int(score1) + 1) / 2
			score = 22 - (6 + distance_turns)
		}
		// append this position to the book
		book = append(book, []int{pos, score})
	}
	return book
}

func To_huffman(b game.Board) int32 {
	// initialize a binary string value of 0
	boardstr := ""
	// run through the board, from left to right, bottom to top: col 0 row 0, col 0 row 1, col 0 row 2...
	for i := 0; i < 7; i++ {
		for j := 0; j < 6; j++ {
			currPos := b.Position() >> j
			currMask := b.Mask() >> j
			if currMask>>(i*7)&1 == 0 {
				boardstr += "0" // column is empty, encode and move on to next column
				break
			}
			if (currMask>>(i*7))&1 == 1 && (currPos>>(i*7))&1 == 1 {
				boardstr += "10" // player 1, encode as 10
			} else {
				boardstr += "11" // player 2, encode as 11
			}
			if j == 5 {
				boardstr += "0" // add 0 to split up columns if column is full
			}
		}
	}
	// after running through the whole board, add one more bit of 0 to fill up full byte (24/32) (32 in our case for 12-ply)
	boardstr += "0"
	// create integer from binary string
	b_int, err := strconv.ParseInt(boardstr, 2, 64)

	if err != nil {
		panic(err)
	}
	// convert to seconds complement if
	if boardstr[2] == 1 && len(boardstr) > 32 {
		b_int -= (2 << 31)
	}
	return int32(b_int)
}

func binary_search(book [][]int, pos int) int {
	l, r := 0, len(book)-1

	for r >= l {
		mid := (l + r + 1) / 2

		p := book[mid][0]
		if p == pos {
			return book[mid][1]
		} else if p < pos {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -121
}

func Get_book_value(b game.Board, p [][]int) int {
	// convert board value to huffman using to_huffman and save it in variable
	huff := To_huffman(b)
	// use binary search to find value in book
	score := binary_search(p, int(huff))
	// if you find the value, return the value found
	if score != -121 {
		return score
	}
	// if you don't find the value, look for the mirrored equivalent
	reverse := game.MakeBoardFromOpening(uint64(reverse(int(b.Position()))), uint64(reverse(int(b.Mask()))), 12)
	huff_rev := To_huffman(reverse)

	score = binary_search(p, int(huff_rev))
	// if you find the value, return the value found
	if score != -121 {
		return score
	}
	// if you don't find the value return 0
	return 0
}

func reverse(bits int) int {
	// algorithm to reverse our board
	res := 0
	for i := 0; i < 7; i++ {
		curr_row := (bits & (0b1111111 << (7 * i))) >> (7 * i) // get current 7 bits. Still need to chop off 7 bits
		res = (res << 7) | curr_row                            // current integers from [7i, 7(i+1)]
	}
	return res
}

// write a function that creates a new db using the entries from the old db using our encoding scheme
// read every entry from the old db and store it in an array
// sort the array to make sure we don't have to spend any time sorting it once we need to read from memory
// one entry at a time write 6 bites for the encoding and 1 byte for the score
