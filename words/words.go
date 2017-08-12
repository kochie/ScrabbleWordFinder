package words

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/kochie/ScrabbleWordFinder/gray"
	"gopkg.in/cheggaaa/pb.v2"
)

var primeMap = make(map[rune]int)
var largestPrime = 1

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// SearchForAnagram will return a list of words in the dictionary that can be made from the letters given.
func SearchForAnagram(letters string, wordTable map[uint64][]string) (possibleWords []string) {
	grayCode := gray.GenerateGrayCode(len(letters))
	productMap := make(map[uint64]bool)
	for _, combo := range grayCode {
		activeLetters := ""
		for i, binary := range combo {
			if 0 == strings.Compare(string(binary), "1") {
				activeLetters += string(letters[i])
			}
		}

		product := getWordProduct(activeLetters)
		if _, ok := productMap[product]; !ok {
			productMap[product] = true
			possibleWords = append(possibleWords, wordTable[product]...)
		}

		for i := 0; i < len(possibleWords); i++ {
			if len(possibleWords[i]) < 3 {
				possibleWords = remove(possibleWords, i)
				i--
				// fmt.Println(tempWords)
			}
		}

		sort.Sort(sort.Reverse(byLength(possibleWords)))

	}
	return
}

// WritePrimeMap writes the result of the primeMap at the current time.
func WritePrimeMap(filename string) {
	data := ""
	for runes, prime := range primeMap {
		fmt.Println(strconv.Itoa(prime))
		data += "char: " + string(runes) + " prime: " + strconv.Itoa(prime) + "\n"
	}
	ioutil.WriteFile(filename, []byte(data), 0644)
	fmt.Println(primeMap)
}

// WriteWordTable will write a wordTable to a file.
func WriteWordTable(wordTable map[uint64][]string, filename string) {
	// data := ""
	bar := pb.StartNew(len(wordTable))
	file, err := os.Create(filename)
	defer file.Close()
	handleError(err)
	writer := bufio.NewWriter(file)
	for product, words := range wordTable {
		fmt.Fprintln(writer, strconv.FormatUint(product, 10)+" - "+strings.Join(words, ", "))
		bar.Increment()
	}
	bar.Finish()
	// ioutil.WriteFile(filename, []byte(data), 0644)
}

// CreateWordTable will create a data object for a list of words.
func CreateWordTable(wordList []string) (wordTable map[uint64][]string) {
	wordTable = make(map[uint64][]string) // make a mapping of int to list of strings.
	bar := pb.StartNew(len(wordList))
	for _, word := range wordList {
		product := getWordProduct(word)
		wordTable[product] = append(wordTable[product], word)
		bar.Increment()
	}
	bar.Finish()
	return
}

// ReadWordList reads a file containing words seperated by a new line into memory.
func ReadWordList(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	handleError(err)
	return strings.Split(string(data), "\n")
}

func getWordProduct(word string) (product uint64) {
	product = 1
	for _, letter := range word {
		product *= uint64(getPrime(letter))
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
			largestPrime = i
			fmt.Println(largestPrime)
			break
		}
	}
	return
}

func testPrimeNumber(i int) bool {
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
