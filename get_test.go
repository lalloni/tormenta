package tormenta_test

import (
	"testing"

	"github.com/jpincas/gouuidv6"
	"github.com/jpincas/tormenta"
	"github.com/jpincas/tormenta/testtypes"
)

func Test_BasicGet(t *testing.T) {
	db, _ := tormenta.OpenTest("data/tests", tormenta.DefaultOptions)
	defer db.Close()

	// Create basic fullStruct and save, then blank the ID
	fullStruct := testtypes.FullStruct{}

	db.Save(&fullStruct)
	ttIDBeforeBlanking := fullStruct.ID
	fullStruct.ID = gouuidv6.UUID{}

	// Attempt to get entity without ID
	found, err := db.Get(&fullStruct)
	if err != nil {
		t.Error("Testing get entity without ID. Got error but should simply fail to find")
	}

	if found {
		t.Errorf("Testing get entity without ID. Expected not to find anything, but did")

	}

	// Reset the fullStruct ID
	fullStruct.ID = ttIDBeforeBlanking
	ok, err := db.Get(&fullStruct)
	if err != nil {
		t.Errorf("Testing basic record get. Got error %v", err)
	}

	if !ok {
		t.Error("Testing basic record get. Record was not found")
	}

}

func Test_GetByID(t *testing.T) {
	db, _ := tormenta.OpenTest("data/tests", tormenta.DefaultOptions)
	defer db.Close()

	fullStruct := testtypes.FullStruct{}
	tt2 := testtypes.FullStruct{}
	db.Save(&fullStruct)

	// Overwite ID
	ok, err := db.Get(&tt2, fullStruct.ID)

	if err != nil {
		t.Errorf("Testing get by id. Got error %v", err)
	}

	if !ok {
		t.Error("Testing get by id. Record was not found")
	}

	if fullStruct.ID != tt2.ID {
		t.Error("Testing get by id. Entity retreived by ID was not the same as that saved")
	}
}

func Test_GetByMultipleIDs(t *testing.T) {
	db, _ := tormenta.OpenTest("data/tests", tormenta.DefaultOptions)
	defer db.Close()

	fullStruct := testtypes.FullStruct{}
	tt2 := testtypes.FullStruct{}
	tt3 := testtypes.FullStruct{}
	toSave := []tormenta.Record{&fullStruct, &tt2, &tt3}
	db.Save(toSave...)

	var results []testtypes.FullStruct
	n, err := db.GetIDs(&results, fullStruct.ID, tt2.ID, tt3.ID)

	if err != nil {
		t.Errorf("Testing get by multiple ids. Got error %v", err)
	}

	if n != 3 {
		t.Errorf("Testing get by multiple ids. Wanted 3 results, got %v", n)
	}

	for i, _ := range results {
		if results[i].ID != toSave[i].GetID() {
			t.Errorf("Testing get by multiple ids. ID mismatch for array member %v. Wanted %v, got %v", i, toSave[i].GetID(), results[i].ID)
		}
	}
}

func Test_GetTriggers(t *testing.T) {
	db, _ := tormenta.OpenTest("data/tests", tormenta.DefaultOptions)
	defer db.Close()

	fullStruct := testtypes.FullStruct{}
	db.Save(&fullStruct)
	ok, err := db.Get(&fullStruct)

	if err != nil {
		t.Errorf("Testing get triggers. Got error %v", err)
	}

	if !ok {
		t.Error("Testing get triggers. Record was not found")
	}

	if !fullStruct.Retrieved {
		t.Error("Testing get triggers.  Expected ttRetrieved = true; got false")
	}
}
