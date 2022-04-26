package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	ErrNotFound = errors.New("not found")
)

type (
	databaseData[T any] struct {
		Data []T `json:"data"`
	}

	Database[T any] interface {
		// Store stores the data inside the json file
		Store(data T) error
		// Get returns the data from the json file given the specified key with the value
		// of the key. If the key is not found, the function returns an error. If more than one items have the key value,
		// the function returns the first it finds.
		Get(key string, value string) (T, error)
		// GetAll returns all the data from the json file
		GetAll() ([]T, error)
		// Delete deletes the data from the json file given the specified key with the value
		// of the key
		Delete(key string, value string) error
		// Update updates the data from the json file given the specified key with the value
		// of the key. If the key is not found, the function returns an error.
		Update(key string, value string, data T) error
	}

	database[T any] struct {
		jsonFile string
	}
)

func NewDatabase[T any](jsonFile string) Database[T] {
	return &database[T]{jsonFile: jsonFile}
}

func (d *database[T]) Store(data T) error {
	dd, err := d.getFileData()
	if err != nil {
		return err
	}

	dd.Data = append(dd.Data, data)

	return d.storeFileData(dd)
}

func (d *database[T]) Get(key string, value string) (T, error) {
	var n T
	dd, err := d.getFileData()
	if err != nil {
		return n, err
	}

	rdata := transformToData(dd.Data)

	for i, d := range rdata {
		for k, v := range d {
			if k == key && fmt.Sprintf("%v", v) == value {
				return dd.Data[i], nil
			}
		}
	}

	return n, ErrNotFound
}

func (d *database[T]) GetAll() ([]T, error) {
	dd, err := d.getFileData()
	if err != nil {
		return nil, err
	}

	var dat []T
	dat = append(dat, dd.Data...)
	return dat, nil
}

func (d *database[T]) Delete(key string, value string) error {
	dd, err := d.getFileData()
	if err != nil {
		return err
	}

	rdata := transformToData(dd.Data)

	var newData []T
	for i := range rdata {
		if fmt.Sprintf("%v", rdata[i][key]) != value {
			newData = append(newData, dd.Data[i])
		}
	}

	if len(newData) == len(dd.Data) {
		return ErrNotFound
	}

	return d.storeFileData(databaseData[T]{Data: newData})
}

func (d *database[T]) Update(key string, value string, data T) error {
	dd, err := d.getFileData()
	if err != nil {
		return err
	}

	var modified bool

	rdata := transformToData(dd.Data)

	for i, d := range rdata {
		for k, v := range d {
			if k == key && fmt.Sprintf("%v", v) == value {
				dd.Data[i] = data
				modified = true
			}
		}
	}

	if !modified {
		return ErrNotFound
	}

	return d.storeFileData(dd)
}

func (d *database[T]) getFileData() (databaseData[T], error) {
	// open json file
	data, err := os.ReadFile(d.jsonFile)
	if err != nil {
		return databaseData[T]{}, err
	}
	var dd databaseData[T]
	err = json.Unmarshal(data, &dd)
	if err != nil {
		return dd, err
	}
	return dd, nil
}

func (d *database[T]) storeFileData(dd databaseData[T]) error {
	// open json file
	data, err := json.Marshal(dd)
	if err != nil {
		return err
	}
	err = os.WriteFile(d.jsonFile, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func transformToData[T any](d []T) []map[string]any {
	data, _ := json.Marshal(d)
	var dataMap []map[string]any
	json.Unmarshal(data, &dataMap)
	return dataMap
}
