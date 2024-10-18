package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot ad word because it already exists")
)

func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]

	if !ok {
		err = ErrNotFound
	}

	return
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
