package maps

type Dictionary map[string]string

func (d Dictionary) Search(word string) (result string) {
	return d[word]
}
