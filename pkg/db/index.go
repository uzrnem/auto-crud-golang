package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	jsonFile = "/app/files/collections.json"
)

type JsonDB struct {
	collections map[string]map[string]map[string]any
}

var (
	DB *JsonDB
)

func Load() error {
	db := &JsonDB{
		collections: map[string]map[string]map[string]any{},
	}
	DB = db
	return DB.ReadFromFile()
}

func (db *JsonDB) WriteToFile() error {
	content, err := json.MarshalIndent(DB.collections, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(jsonFile, content, 0644)
}

func (db *JsonDB) ReadFromFile() error {
	content, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		if strings.Contains(err.Error(), "no such file") {
			return nil
		}
		return err
	}
	return json.Unmarshal(content, &DB.collections)
}

func (db *JsonDB) SetupCollection(collection string) {
	DB.collections[collection] = map[string]map[string]any{}
}

func (db *JsonDB) GetDocumentById(collection, id string) (map[string]any, error) {
	err := db.ReadFromFile()
	if err != nil {
		return nil, err
	}
	data := DB.collections[collection][id]
	if DB.collections[collection] == nil || data == nil {
		return nil, errors.New("document not found")
	}
	return data, nil
}

func (db *JsonDB) DeleteDocumentById(collection, id string) error {
	err := db.ReadFromFile()
	if err != nil {
		return err
	}
	data := DB.collections[collection][id]
	if DB.collections[collection] == nil || data == nil {
		return errors.New("document not found")
	}
	if DB.collections[collection] != nil {
		delete(DB.collections[collection], id)
		return db.WriteToFile()
	}

	return nil
}

func (db *JsonDB) GetDocuments(collection string) (map[string]map[string]any, error) {
	err := db.ReadFromFile()
	if err != nil {
		return nil, err
	}
	if DB.collections[collection] == nil {
		return map[string]map[string]any{}, nil
	}
	return DB.collections[collection], nil
}

func (db *JsonDB) SaveDocument(collection string, data map[string]any) error {
	if DB.collections[collection] == nil {
		DB.collections[collection] = map[string]map[string]any{}
	}
	DB.collections[collection][fmt.Sprint(data["id"])] = data
	return db.WriteToFile()
}
