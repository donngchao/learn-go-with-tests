package main

// Dictionary store definitions to words
type Dictionary map[string]string

// Search find a word in the dictionary
func (d Dictionary) Search(word string) string {
	return d[word]
}

type Telephone map[string]int

func (t Telephone) Search(name string) int {
	return t[name]
}

type Price map[int]int

func (p Price) GetPrice(tag int) int {
	return p[tag]
}
