package maps

import (
	"go-with-tdd/coverage"
	"testing"
)

func TestMain(t *testing.M) {
	coverage.EnforceCoverage(t, 1, "maps")
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("should return a word for known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is just a test"
		assertString(t, got, expected)
	})

	t.Run("should return an error for unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("should add a new key value", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, value)
		assertNoError(t, err)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("should not add value over an existing key", func(t *testing.T) {
		key := "test"
		value := "this is a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new value")
		assertError(t, err, ErrKeyExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should update an existing key's value", func(t *testing.T) {
		key := "test"
		value := "this is a test"
		dictionary := Dictionary{key: value}
		newValue := "updated test"

		err := dictionary.Update(key, newValue)
		assertNoError(t, err)
		assertDefinition(t, dictionary, key, newValue)
	})

	t.Run("should throw if try to update a new key value", func(t *testing.T) {
		key := "test"
		value := "this is test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)
		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "test value"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", key)
	}
}

func assertString(t *testing.T, got string, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("got %q but expected %q when given %q", got, expected, "test")
	}
}

func assertError(t *testing.T, got, expected error) {
	t.Helper()
	if got != expected {
		t.Errorf("got error %q but expected %q", got, expected)
	}

	if got == nil {
		if expected == nil {
			return
		}
		t.Fatal("expected to get an error")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should get the newly added word")
	}

	if got != value {
		t.Errorf("got %q but expected %q when given %q", got, value, "test")
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got an error but didn't expected")
	}
}
