package solver

import (
	"io"
	"math/big"
	"os"
	"strconv"
	"sync"

	"github.com/jacobomantilla10/connect-four/game"
)

// create openingBook which is the same type as transposition
type OpeningBook struct {
	Openings [][]int
	Size     int
}

func MakeOpeningBook() OpeningBook {
	filename := "../solver/bookDeepDist.dat"

	book := read_book(filename)
	//book := read_book(filename)
	return OpeningBook{Openings: book, Size: len(book)}
}

func read_book(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// get file length
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	numGoroutines := 56                        // Got this number through trial and error
	length := int(fileInfo.Size())             // length of file in bytes
	numItems := length / 5                     // number of items. 5 bytes give us one "opening" containing position and score.
	chunkSize := (length - 15) / numGoroutines // size of chunks in bytes after we split file up into numGoroutines goroutines
	//chunkSize := (length) / numGoroutines
	book := make([][]int, numItems) // book that will hold the openings

	// We have a file of 21004495 bytes we will read with 56 goroutines, each reading it by chunks of 5 bytes at a time.
	// To be able to do this we need to read the last 15 bytes, because 21004495-15 is divisible by 56 and then again by 5.
	pos, score := read_line(file, length-15, 0)
	book[len(book)-3] = []int{pos, score}
	pos, score = read_line(file, length-10, 0)
	book[len(book)-2] = []int{pos, score}
	pos, score = read_line(file, length-5, 0)
	book[len(book)-1] = []int{pos, score}

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Spin off numGoroutines goroutines to read the file by chunks
	for i := 0; i < numGoroutines; i++ {
		go func(i, offset int) {
			defer wg.Done()
			k := 0
			for j := 0; j < chunkSize; j += 5 {
				pos, score := read_line(file, offset, j)
				idx := (i * (chunkSize / 5)) + k
				book[idx] = []int{pos, score}
				k++
			}
		}(i, (i * chunkSize))
	}
	wg.Wait()
	return book
}

func read_line(file *os.File, offset, j int) (int, int) {
	posBuf := make([]byte, 4)   // first four bytes give us the position
	scoreBuf := make([]byte, 1) // last byte gives us the score of the position

	// Read 4 bytes from file. These give us our position in the Huffman encoded format.
	_, err := file.ReadAt(posBuf, int64(offset+j))
	if err != nil && err != io.EOF {
		panic(err)
	}
	pos := int(big.NewInt(0).SetBytes(posBuf).Int64())
	// Check for overflow and correct. TODO is this correct?
	if pos > 0x7FFFFFFF {
		pos -= (2 << 31)
	}
	// read one more byte into buffer2, this represents the score.
	file.ReadAt(scoreBuf, int64(offset+j+4))
	score1 := int8(scoreBuf[0])

	// Convert score from db format into format used by engine. TODO add more here about formats.
	var score int
	if score1 < 0 {
		distance_turns := (int(score1) + 100 + 1) / 2
		score = (22 - (6 + distance_turns)) * -1
	} else if score1 > 0 {
		distance_turns := (100 - int(score1) + 1) / 2
		score = 22 - (6 + distance_turns)
	}

	return pos, score
}

func (ob OpeningBook) Get_book_value(b game.Board) int {
	p := ob.Openings
	// convert board value to huffman using to_huffman and save it in variable
	huff := to_huffman(b)
	// use binary search to find value in book
	score := binary_search(p, int(huff))
	// if you find the value, return the value found
	if score != -121 {
		return score
	}
	// if you don't find the value, look for the mirrored equivalent
	reverse := game.MakeBoardFromOpening(uint64(reverse(int(b.Position()))), uint64(reverse(int(b.Mask()))), 12)
	huff_rev := to_huffman(reverse)

	score = binary_search(p, int(huff_rev))
	// if you find the value, return the value found
	if score != -121 {
		return score
	}
	// if you don't find the value return 0
	return -300
}

func to_huffman(b game.Board) int32 {
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
	b_int, err := strconv.ParseInt(boardstr, 2, 32)

	if err != nil {
		b_int, _ = strconv.ParseInt(boardstr, 2, 64)
		b_int -= (2 << 31)
	}
	// convert to seconds complement if necessary
	// if boardstr[2] == 1 && len(boardstr) > 32 {
	// 	b_int -= (2 << 31)
	// }
	return int32(b_int)
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
