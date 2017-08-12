package main

import (
	"github.com/kochie/ScrabbleWordFinder/words"
)

func main() {
	words.WritePrimeMap("prime_a.txt")
	words.WriteWordTable(words.CreateWordTable(words.ReadWordList("./data/words.txt")), "./data/wordTable.txt")
	words.WritePrimeMap("prime_b.txt")

}
