package english_bible_test

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/brianglass/english_bible"
	_ "github.com/mattn/go-sqlite3"
)

func TestScriptureLookup(t *testing.T) {
	db, e := sql.Open("sqlite3", "bible.db")
	if e != nil {
		t.Errorf("Got error opening database: %#v.", e)
	}
	bible := english_bible.NewBible(db)

	testCases := []struct {
		reference string
		count     int
	}{
		{"Matt 1.1-25", 25},
		{"Matt 4.25-5.13", 14},
		{"Matt 10.32-36, 11.1", 6},
		{"Matt 6.31-34, 7.9-11", 7},
		{"Matt 10.1, 5-8", 5},
		{"Mark 15.22, 25, 33-41", 11},
		// single chapter book
		{"Jude 1-10", 10},
		{"1 John 2.7-17", 11},
		{"Gen 17.1-2, 4, 5-7, 8, 9-10, 11-12, 14", 12},
		// discontinuous chapters
		{"Job 38.1-23; 42.1-5", 28},
		// multiple books
		{"1 Cor 5.6-8; Gal 3.13-14", 5},
		// multiple books with : instead of .
		{"Matt 26:2-20; John 13:3-17; Matt 26:21-39; Luke 22:43-45; Matt 26:40-27:2", 94},
		// individual full chapters
		{"Prov 10, 3, 8", 32 + 35 + 36},
		// Multiple chapters
		{"Jonah 1.1-4.11", 17 + 10 + 10 + 11},
		// Deuterocanonical
		{"4 Kgs 2.6-14", 9},
		{"Baruch 3.35-4.4", 3 + 4},
		{"Wis 3.1-9", 9},
	}

	for _, tc := range testCases {
		t.Run("Scripture Lookup", func(t *testing.T) {
			passage := bible.Lookup(tc.reference)
			// Not really a rigorous test, but it ought to catch a regression ;)
			if len(passage) != tc.count {
				actual, _ := json.MarshalIndent(passage, "", "\t")
				t.Errorf("%s", actual)
				t.Errorf("%s should return %d verses but returned %d verses.", tc.reference, tc.count, len(passage))
				t.Fail()
			}
		})
	}
}
