package main

type Dictionary map[string]string

func (d *Dictionary) Search(entry string) string {
	return (*d)[entry]
}
