package main

func emptyBoard() [7][6]string {
	return [7][6]string{
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
	}
}

func boardFromString(arr [7][6]string, board string) [7][6]string {
	for i, ch := range board {
		// convert the current utf character into a number and subtract one to find the column we need to insert into
		// loop into first open spot in array and add the correct value depending on parity of turn
		insertCol := int(ch-'0') - 1
		insertRow := 0
		insertVal := "Player1"
		if i%2 == 1 {
			insertVal = "Player2"
		}
		for j := 0; arr[insertCol][j] != "Empty" && j < 6; j++ {
			insertRow++
		}
		if insertRow != 6 {
			arr[insertCol][insertRow] = insertVal
		}
	}
	return arr
}

func newTemplateData() templateData {
	return templateData{
		BoardString: "",
		Board:       emptyBoard(),
		IsGameOver:  false,
	}
}
