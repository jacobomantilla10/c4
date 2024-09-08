package solver

import "testing"

func BenchmarkRead_book(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeOpeningBook()
	}
}

func TestRead_book(t *testing.T) {
	t.Run("Make opening book", func(t *testing.T) {
		MakeOpeningBook()
	})
}
