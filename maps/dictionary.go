package main

import "errors"

type Dictionary map[string]string

func (d *Dictionary) Search(entry string) (string, error) {
	if definition, ok := (*d)[entry]; ok {
		return definition, nil
	} else {
		return "", errors.New("word unknown")
	}
}
