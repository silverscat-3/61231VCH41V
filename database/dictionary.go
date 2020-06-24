package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	dictionary struct {
		ID      int
		String1 string
		String2 string
		String3 string
	}
	DictionaryModel struct {
		DB *sqlx.DB
	}
)

func (dm *DictionaryModel) DeleteOldData() error {
	time := time.Now()
	time = time.AddDate(0, -1, 0)

	if _, err := dm.DB.NamedExec("DELETE dictionary id < :timestamp", time.Unix()); nil != err {
		return err
	}

	return nil
}

func (dm *DictionaryModel) Get() ([][]string, error) {
	dics := []*dictionary{}
	if err := dm.DB.Select(&dics, "SELECT * FROM dictionary"); nil != err {
		return nil, err
	}

	array := [][]string{}
	for _, dic := range dics {
		array = append(array, []string{dic.String1, dic.String2, dic.String3})
	}

	return array, nil
}

func (dm *DictionaryModel) Put(array [][]string) error {
	for _, words := range array {
		if _, err := dm.DB.NamedExec("INSERT INTO dictionary (id, string1, string2, string3) VALUES(:id, :string1, :string2, :string3)", map[string]interface{}{"id": time.Now().UnixNano(), "string1": words[0], "string2": words[1], "string3": words[2]}); nil != err {
			return err
		}
	}

	return nil
}
