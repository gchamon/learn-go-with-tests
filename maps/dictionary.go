package main

type Dictionary map[string]string

func (d *Dictionary) Search(entry string) (string, error) {
	return (*d)[entry], nil
}
