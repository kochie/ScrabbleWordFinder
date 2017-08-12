package words

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var primeMap = make(map[rune]int)
var largestPrime = 1

// WritePrimeMap writes the result of the primeMap at the current time.
func WritePrimeMap(filename string) {
	data := ""
	for runes, prime := range primeMap {
		fmt.Println(prime)
		data += string(runes) + " prime: " + string(prime) + "\n"
	}
	ioutil.WriteFile(filename, []byte(data), 0644)
}

// WriteWordTable will write a wordTable to a file.
func WriteWordTable(wordTable map[int][]string, filename string) {
	data := ""
	for product, words := range wordTable {
		data += string(product) + strings.Join(words, ", ") + "\n"
	}
	ioutil.WriteFile(filename, []byte(data), 0644)
}

// CreateWordTable will create a data object for a list of words.
func CreateWordTable(wordList []string) (wordTable map[int][]string) {
	wordTable = make(map[int][]string) // make a mapping of int to list of strings.

	for _, word := range wordList {
		product := getWordProduct(word)
		wordTable[product] = append(wordTable[product], word)
	}
	return
}

// ReadWordList reads a file containing words seperated by a new line into memory.
func ReadWordList(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	handleError(err)
	return strings.Split(string(data), "\n")
}

func getWordProduct(word string) (product int) {
	product = 1
	for _, letter := range word {
		product *= getPrime(letter)
	}
	return
}

func getPrime(letter rune) (prime int) {
	if val, ok := primeMap[letter]; ok {
		//do something here
		prime = val
	} else {
		prime = generatePrime(largestPrime)
		primeMap[letter] = prime
	}
	return
}

func generatePrime(previousPrime int) (prime int) {
	i := previousPrime
	for {
		i++
		if testPrimeNumber(i) {
			prime = i
			break
		}
	}
	return
}

func testPrimeNumber(i int) bool {
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return true
		}
	}
	return false
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
