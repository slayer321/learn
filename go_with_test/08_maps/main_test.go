package main

import (
	"testing"
)

// func TestSearch(t *testing.T) {
// 	dictionary := Dictionary{"test": "this is just a test"}
// 	got := dictionary.Search("sachin")

// 	want := "this is just a test"
// 	assertStrings(t, got, want)
// }

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

// func TestSearch(t *testing.T) {
// 	dictionary := Dictionary{"test": "this is just a test"}

// 	t.Run("know word", func(t *testing.T) {
// 		got, _ := dictionary.Search("test")
// 		want := "this is just a test"
// 		assertStrings(t, got, want)
// 	})

// 	t.Run("unkown word", func(t *testing.T) {
// 		_, err := dictionary.Search("unknow")

// 		if err == nil {
// 			t.Fatal("expected to get an error.")
// 		}
// 		assertError(t, err, ErrNotFound)
// 	})
// }

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("update test", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		newDef := "new Def"

		err := dictionary.Update(word, newDef)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDef)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete", func(t *testing.T) {
		word := "test"
		def := "this is just test"
		dic := Dictionary{word: def}

		dic.Delete(word)
		_, err := dic.Search(word)
		if err != ErrNotFound {
			t.Errorf("Expected %q to be deleted", word)
		}

	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err.Error())
	}
	if definition != got {
		t.Errorf("got error %q want %q", got, definition)
	}
}
