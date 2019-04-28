package dictionary

import "errors"

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

type Dictionary map[string]string

var ErrNotFound = errors.New("word not found")
var ErrExists = errors.New("word already exists")

func (d Dictionary) Search(word string) (string, error) {
	val, wasFound := d[word]
	if !wasFound {
		return "", ErrNotFound
	}
	return val, nil
}

// NOTE
// - this is one of the common function patterns in go, where you return
//   `(value, error)`

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrExists
	default:
		return err
	}

	return nil
}

// NOTE
// - map is a reference type - we never wind up making copies of it

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
