package dict

import "errors"
import "fmt"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not found")
	errCantUpdate = errors.New("Can't update non-exisitng word")
	errCantDelete = errors.New("Can't delete non-exisitng word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	//switch err {
	//case errNotFound:
	//case nil:
	//  return errWordExists
	//}

	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil

}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil

}

// Delete a word
func (d Dictionary) Delete(word string) {

	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		fmt.Println(errCantDelete)
	}
}
