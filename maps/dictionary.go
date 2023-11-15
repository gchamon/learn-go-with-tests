package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("word unknown")

func (d *Dictionary) Search(entry string) (string, error) {
	if definition, ok := (*d)[entry]; ok {
		return definition, nil
	} else {
		return "", ErrNotFound
	}
}

func (d *Dictionary) Add(entry, definition string) {

}
