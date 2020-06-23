package markov

import (
	"log"
	"math/rand"
	"time"

	"github.com/bluele/mecab-golang"
)

func ParseToNode(m *mecab.MeCab, input string) []string {
	words := []string{}

	tg, err := m.NewTagger()
	if nil != err {
		log.Fatalln(err)
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(input)
	if nil != err {
		log.Fatalln(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		if "" != node.Surface() {
			words = append(words, node.Surface())
		}
		if nil != node.Next() {
			break
		}
	}

	return words
}

func GetMarkovBlocks(words []string) [][]string {
	res := [][]string{}
	resHead := []string{}
	resEnd := []string{}

	if 3 > len(words) {
		return res
	}

	resHead = []string{"#*#* BOS *#*#", words[0], words[1]}
	res = append(res, resHead)

	for i := 1; i < len(words)-2; i++ {
		markovBlock := []string{words[i], words[i+1], words[i+2]}
		res = append(res, markovBlock)
	}

	resEnd = []string{words[len(words)-2], words[len(words)-1], "#*#* EOS *#*#"}
	res = append(res, resEnd)

	return res
}

func findBlocks(array [][]string, target string) [][]string {
	blocks := [][]string{}
	for _, s := range array {
		if s[0] == target {
			blocks = append(blocks, s)
		}
	}

	return blocks
}

func connectBlocks(array [][]string, dist []string) []string {
	rand.Seed(time.Now().UnixNano())
	i := 0

	for _, word := range array[rand.Intn(len(array))] {
		if 0 != i {
			dist = append(dist, word)
		}
		i++
	}

	return dist
}

func MarkovChainExec(array [][]string) []string {
	ret := []string{}
	block := [][]string{}
	count := 0

	block = findBlocks(array, "#*#* BOS *#*#")
	ret = connectBlocks(block, ret)

	for "#*#* EOS *#*#" != ret[len(ret)-1] {
		block = findBlocks(array, ret[len(ret)-1])
		if 0 == len(block) {
			break
		}
		ret = connectBlocks(block, ret)

		count++
		if 128 == count {
			break
		}
	}

	return ret
}
