package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExitsts = errors.New("Already exists")
var errCantUpdate = errors.New("No matches, it can't be updated")
var errCantDelete = errors.New("No matches, it can't be deleted")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExitsts
	}
	return nil
}

// Updated a word's definition in the dictionary
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word and the definition from the dictionary
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		// delete the word from dictionary
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}

