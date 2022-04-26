package database

import (
	"os"
	"testing"
)

type TestData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Valid bool   `json:"valid"`
}

func reset() {
	os.Remove("/tmp/test_db.json")
	os.WriteFile("/tmp/test_db.json", []byte("{}"), 0644)
}

func TestMain(m *testing.M) {
	// create db on tmp/test_db.json
	os.WriteFile("/tmp/test_db.json", []byte("{}"), 0644)
	m.Run()
	// remove db
	os.Remove("/tmp/test_db.json")
}

func getDB[T any]() Database[T] {
	return NewDatabase[T]("/tmp/test_db.json")
}

// Test Store
func TestStoreAndGet(t *testing.T) {
	reset()
	db := getDB[TestData]()
	td := TestData{ID: 1, Name: "test", Age: 10, Valid: true}
	err := db.Store(td)
	if err != nil {
		t.Errorf("Error storing data: %s", err)
	}

	tdr, err := db.Get("id", "1")
	if err != nil {
		t.Errorf("Error getting data: %s", err)
	}

	if tdr.ID != td.ID {
		t.Errorf("ID does not match: %d != %d", tdr.ID, td.ID)
	}

	if tdr.Name != td.Name {
		t.Errorf("Name does not match: %s != %s", tdr.Name, td.Name)
	}

	if tdr.Age != td.Age {
		t.Errorf("Age does not match: %d != %d", tdr.Age, td.Age)
	}

	if tdr.Valid != td.Valid {
		t.Errorf("Valid does not match: %t != %t", tdr.Valid, td.Valid)
	}
}

// Test GetAll
func TestGetAll(t *testing.T) {
	reset()
	db := getDB[TestData]()
	td := TestData{ID: 1, Name: "test", Age: 10, Valid: true}
	td2 := TestData{ID: 2, Name: "test2", Age: 20, Valid: true}
	err := db.Store(td)
	if err != nil {
		t.Errorf("Error storing data: %s", err)
	}

	err = db.Store(td2)
	if err != nil {
		t.Errorf("Error storing data: %s", err)
	}

	tds, err := db.GetAll()
	if err != nil {
		t.Errorf("Error getting data: %s", err)
	}

	if len(tds) != 2 {
		t.Errorf("Expected 2 data, got %d", len(tds))
	}

	if tds[0].ID != td.ID {
		t.Errorf("ID does not match: %d != %d", tds[0].ID, td.ID)
	}

	if tds[0].Name != td.Name {
		t.Errorf("Name does not match: %s != %s", tds[0].Name, td.Name)
	}

	if tds[0].Age != td.Age {
		t.Errorf("Age does not match: %d != %d", tds[0].Age, td.Age)
	}

	if tds[0].Valid != td.Valid {
		t.Errorf("Valid does not match: %t != %t", tds[0].Valid, td.Valid)
	}
}

// Test Update
func TestUpdate(t *testing.T) {
	reset()
	db := getDB[TestData]()
	td := TestData{ID: 1, Name: "test", Age: 10, Valid: true}
	err := db.Store(td)
	if err != nil {
		t.Errorf("Error storing data: %s", err)
	}

	td.Name = "test2"
	err = db.Update("id", "1", td)
	if err != nil {
		t.Errorf("Error updating data: %s", err)
	}

	tdr, err := db.Get("id", "1")
	if err != nil {
		t.Errorf("Error getting data: %s", err)
	}

	if tdr.ID != td.ID {
		t.Errorf("ID does not match: %d != %d", tdr.ID, td.ID)
	}

	if tdr.Name != td.Name {
		t.Errorf("Name does not match: %s != %s", tdr.Name, td.Name)
	}

	if tdr.Age != td.Age {
		t.Errorf("Age does not match: %d != %d", tdr.Age, td.Age)
	}

	if tdr.Valid != td.Valid {
		t.Errorf("Valid does not match: %t != %t", tdr.Valid, td.Valid)
	}
}

// Test Delete
func TestDelete(t *testing.T) {
	reset()
	db := getDB[TestData]()
	td := TestData{ID: 1, Name: "test", Age: 10, Valid: true}
	err := db.Store(td)
	if err != nil {
		t.Errorf("Error storing data: %s", err)
	}

	td2 := TestData{ID: 2, Name: "test2", Age: 20, Valid: true}
	err = db.Store(td2)
	if err != nil {
		t.Errorf("Error storing data: %s", err)
	}

	err = db.Delete("id", "1")
	if err != nil {
		t.Errorf("Error deleting data: %s", err)
	}

	_, err = db.Get("id", "1")
	if err == nil {
		t.Errorf("Expected error getting data, got nil")
	}

	tds, err := db.GetAll()
	if err != nil {
		t.Errorf("Error getting data: %s", err)
	}

	if len(tds) != 1 {
		t.Errorf("Expected 1 data, got %d", len(tds))
	}
}

func TestGetNotFound(t *testing.T) {
	reset()
	db := getDB[TestData]()
	_, err := db.Get("id", "999")
	if err == nil {
		t.Errorf("Expected error getting data, got nil")
	}
}

func TestUpdateNotFound(t *testing.T) {
	reset()
	db := getDB[TestData]()
	td := TestData{ID: 1, Name: "test", Age: 10, Valid: true}
	err := db.Update("id", "999", td)
	if err == nil {
		t.Errorf("Expected error updating data, got nil")
	}
}

func TestDeleteNotFound(t *testing.T) {
	reset()
	db := getDB[TestData]()
	err := db.Delete("id", "999")
	if err == nil {
		t.Errorf("Expected error deleting data, got nil")
	}
}
