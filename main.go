package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kochie/ScrabbleWordFinder/words"
)

func main() {
	// words.WritePrimeMap("prime_a.txt")
	wordTable := words.CreateWordTable(words.ReadWordList("./data/words.txt"))
	words.WriteWordTable(wordTable, "./data/wordTable.txt")
	words.WritePrimeMap("prime_b.txt")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(words.SearchForAnagram(text, wordTable))
	}

}
