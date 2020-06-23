package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bluele/mecab-golang"
	"github.com/silverscat-3/61231VCH41V/markov"
	"github.com/silverscat-3/61231VCH41V/misskey"
)

func main() {
	if 2 > len(os.Args) {
		printHelp()
	}

	switch os.Args[1] {
	case "post":
		if 4 > len(os.Args) {
			log.Fatalln("Too few arguments.")
		}

		host := os.Args[2]
		token := os.Args[3]

		postNote(host, token)

	case "help":
		printHelp()

	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		printHelp()
	}
}

func postNote(host, token string) {
	m, err := misskey.NewMisskey(host, token)
	if nil != err {
		log.Fatalln(err)
	}

	note := genString()

	if err := m.NotePost(note); nil != err {
		log.Fatalln(err)
	}
	log.Printf("Posted!! %s\n", trimString(note))
}

func printHelp() {
	fmt.Print(`
61231VCH41V is shitpost maker.

Usage:

        61231VCH41V <command> [arguments]

The commands are:
post	post a shitpost
`)
	os.Exit(0)
}

func genString() string {
	notes := loadFile("output.txt")

	m, err := mecab.New()
	if nil != err {
		log.Fatalln(err)
	}
	defer m.Destroy()

	words := [][]string{}
	for _, note := range notes {
		words = append(words, markov.ParseToNode(m, note))
	}

	markovBlocks := [][]string{}
	for _, word := range words {
		markovBlocks = append(markovBlocks, markov.GetMarkovBlocks(word)...)
	}
	result1 := markov.MarkovChainExec(markovBlocks)
	result2 := markov.MarkovChainExec(markovBlocks)
	result3 := markov.MarkovChainExec(markovBlocks)
	result := []string{strings.Join(result1, ""), strings.Join(result2, ""), strings.Join(result3, "")}

	return strings.Join(result, "")
}

func trimString(s string) string {
	result := []rune{}
	for i, r := range []rune(s) {
		result = append(result, r)
		if 31 < i {
			result = append(result, ' ')
			result = append(result, '.')
			result = append(result, '.')
			result = append(result, '.')

			break
		}
	}

	return string(result)
}

func loadFile(fileName string) []string {
	f, err := os.Open(fileName)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	list := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	return list
}
