package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bluele/mecab-golang"
	"github.com/jmoiron/sqlx"
	"github.com/silverscat-3/61231VCH41V/database"
	"github.com/silverscat-3/61231VCH41V/markov"
	"github.com/silverscat-3/61231VCH41V/misskey"
)

func main() {
	if 2 > len(os.Args) {
		printHelp()
	}

	switch os.Args[1] {
	case "post":
		if 5 > len(os.Args) {
			log.Fatalln("Too few arguments. 4 arguments need.")
		}

		host := os.Args[2]
		token := os.Args[3]
		dbPath := os.Args[4]
		db, err := database.Connect(dbPath)
		if nil != err {
			log.Fatalln(err)
		}

		postNote(host, token, db)

	case "learning":
		if 3 > len(os.Args) {
			log.Fatalln("Too few arguments. 2 arguments need.")
		}

		text := os.Args[2]
		dbPath := os.Args[3]
		db, err := database.Connect(dbPath)
		if nil != err {
			log.Fatalln(err)
		}

		learning(db, text)

	case "help":
		printHelp()

	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		printHelp()
	}
}

func postNote(host, token string, db *sqlx.DB) {
	m, err := misskey.NewMisskey(host, token)
	if nil != err {
		log.Fatalln(err)
	}

	note, err := genString(db)
	if nil != err {
		log.Fatalln(err)
	}

	if err := m.NotePost(note); nil != err {
		log.Fatalln(err)
	}

	log.Printf("Posted!! %s\n", trimString(note))
}

func learning(db *sqlx.DB, text string) {
	m, err := mecab.New()
	if nil != err {
		log.Fatalln(err)
	}
	defer m.Destroy()

	words := markov.ParseToNode(m, text)
	blocks := markov.GetMarkovBlocks(words)

	dm := &database.DictionaryModel{db}
	dm.Put(blocks)

	log.Printf("Learned! %s", words)
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

func genString(db *sqlx.DB) (string, error) {
	dm := database.DictionaryModel{db}
	markovBlocks, err := dm.Get()
	if nil != err {
		return "", err
	}

	result1 := markov.MarkovChainExec(markovBlocks)
	result2 := markov.MarkovChainExec(markovBlocks)
	result3 := markov.MarkovChainExec(markovBlocks)
	result := []string{strings.Join(result1, ""), strings.Join(result2, ""), strings.Join(result3, "")}

	return strings.Join(result, ""), nil
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
