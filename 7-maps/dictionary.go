package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]

	if !ok {
		err = ErrNotFound
	}

	return
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
