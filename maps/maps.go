package maps

import "github.com/pkg/errors"

type Dictionary map[string]string

var (
	ErrNotFound  = errors.Wrap(errors.New("Key not known to dictionary"), "Error")
	ErrKeyExists = errors.Wrap(errors.New("Key already exists"), "Error")
)

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return word, nil
}

func (d Dictionary) Add(key, value string) (returnedError error) {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
		returnedError = nil
	case nil:
		returnedError = ErrKeyExists
	}

	return
}

func (d Dictionary) Update(key, newValue string) (returnedError error) {
	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = newValue
		returnedError = nil
	case ErrNotFound:
		returnedError = ErrNotFound
	}

	return
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
