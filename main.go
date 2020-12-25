package main

import (
	"fmt"

	"github.com/shortdaddy0711/learngo/mydict"

)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	word := "hello"
	definition1 := "Greeting"
	definition2 := "world"

	err := dictionary.Add(word, definition1)
	if err != nil {
		fmt.Println(err)
	}

	def, err := dictionary.Search(word)
	fmt.Println("Found", word, "the definition:", def)

	err2 := dictionary.Update(word, definition2)
	if err2 != nil {
		fmt.Println(err2)
	}

	def2, err := dictionary.Search(word)
	fmt.Println("Found", word, "the definition:", def2)

	err3 := dictionary.Delete(word)
	if err3 != nil {
		fmt.Println(err3)
	}

	_, err4 := dictionary.Search(word)
	if err4 != nil {
		fmt.Println(err4)
	}
}