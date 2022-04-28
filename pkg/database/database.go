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
	databaseData struct {
		Data []interface{} `json:"data"`
	}

	Database interface {
		// Store stores the data inside the json file
		Store(data interface{}) error
		// Get returns the data from the json file given the specified key with the value
		// of the key. If the key is not found, the function returns an error. If more than one items have the key value,
		// the function returns the first it finds.
		Get(key string, value string, data interface{}) error
		// GetAll returns all the data from the json file
		GetAll(data interface{}) error
		// Delete deletes the data from the json file given the specified key with the value
		// of the key
		Delete(key string, value string) error
		// Update updates the data from the json file given the specified key with the value
		// of the key. If the key is not found, the function returns an error.
		Update(key string, value string, data interface{}) error
	}

	database struct {
		jsonFile string
	}
)

func NewDatabase(jsonFile string) Database {
	return &database{jsonFile: jsonFile}
}

func (d *database) Store(data interface{}) error {
	dd, err := d.getFileData()
	if err != nil {
		return err
	}

	dd.Data = append(dd.Data, data)

	return d.storeFileData(dd)
}

func (d *database) Get(key string, value string, data interface{}) error {
	dd, err := d.getFileData()
	if err != nil {
		return err
	}

	rdata := transformToData(dd.Data)

	for i, d := range rdata {
		for k, v := range d {
			if k == key && fmt.Sprintf("%v", v) == value {
				md := dd.Data[i]
				d, _ := json.Marshal(md)
				return json.Unmarshal(d, &data)
			}
		}
	}

	return ErrNotFound
}

func (d *database) GetAll(data interface{}) error {
	dd, err := os.ReadFile(d.jsonFile)
	if err != nil {
		return err
	}

	var dbd databaseData

	err = json.Unmarshal(dd, &dbd)
	if err != nil {
		return err
	}

	dbdd, _ := json.Marshal(dbd.Data)
	return json.Unmarshal(dbdd, &data)
}

func (d *database) Delete(key string, value string) error {
	dd, err := d.getFileData()
	if err != nil {
		return err
	}

	rdata := transformToData(dd.Data)

	var newData []interface{}
	for i := range rdata {
		if fmt.Sprintf("%v", rdata[i][key]) != value {
			newData = append(newData, dd.Data[i])
		}
	}

	if len(newData) == len(dd.Data) {
		return ErrNotFound
	}

	return d.storeFileData(databaseData{Data: newData})
}

func (d *database) Update(key string, value string, data interface{}) error {
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

func (d *database) getFileData() (databaseData, error) {
	// open json file
	data, err := os.ReadFile(d.jsonFile)
	if err != nil {
		return databaseData{}, err
	}
	var dd databaseData
	err = json.Unmarshal(data, &dd)
	if err != nil {
		return dd, err
	}
	return dd, nil
}

func (d *database) storeFileData(dd databaseData) error {
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

func transformToData(d []interface{}) []map[string]interface{} {
	data, _ := json.Marshal(d)
	var dataMap []map[string]interface{}
	json.Unmarshal(data, &dataMap)
	return dataMap
}
