package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bluele/mecab-golang"
	"github.com/silverscat-3/612311VCH41V/markov"
)

func main() {
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
	result1[len(result1)-1] = ""
	if "。" != result1[len(result1)-1] {
		result1 = append(result1, "。")
	}

	result2 := markov.MarkovChainExec(markovBlocks)
	result2[len(result2)-1] = ""
	if "。" != result2[len(result2)-1] {
		result2 = append(result2, "。")
	}

	result3 := markov.MarkovChainExec(markovBlocks)
	result3[len(result3)-1] = ""
	if "。" != result3[len(result3)-1] {
		result3 = append(result3, "。")
	}

	result := []string{strings.Join(result1, ""), strings.Join(result2, ""), strings.Join(result3, "")}

	fmt.Println(strings.Join(result, ""))
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