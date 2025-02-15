package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{"Success", "Pokhrel", "9876543120"})
	data = append(data, Entry{"Biththal", "Acharya", "1234567890"})
	data = append(data, Entry{"John", "Black", "2103489449"})

	//Differentiate between the commands
	switch arguments[1] {
	//The search command
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)

	// The list command
	case "list":
		list()
	// Response to anything that is not a match
	default:
		fmt.Println("Not a valid option")
	}
}
