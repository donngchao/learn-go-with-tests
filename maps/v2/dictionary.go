package main

import "errors"

// Dictionary store definitions to words
type Dictionary map[string]string
type PhoneBook map[string]int
// ErrNotFound means the definition could not be found for the given word
var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrNoSuchPerson = errors.New("could not find this guy in the PhoneBook")
// Search find a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (p PhoneBook) Search(people string)(int,error)  {
	phoneNum, ok := p[people]
	if !ok{
		return 0,ErrNoSuchPerson
	}
	return phoneNum,nil
}
