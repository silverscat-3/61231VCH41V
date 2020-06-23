package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/bluele/mecab-golang"
	"github.com/silverscat-3/61231VCH41V/markov"
	"github.com/silverscat-3/61231VCH41V/misskey"
	flag "github.com/spf13/pflag"
)

func main() {
	var (
		misskeyToken = flag.String("misskey_token", "", "API token of Misskey.")
	)

	flag.Parse()

	m, err := misskey.NewMisskey("https://groundpolis.app/", *misskeyToken)
	if nil != err {
		log.Fatalln(err)
	}

	if err := m.NotePost(genString()); nil != err {
		log.Fatalln(err)
	}
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
