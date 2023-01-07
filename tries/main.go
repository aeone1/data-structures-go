package main

import (
	"fmt"
	"github.com/aeone1/data-structures-go/tries/array-trie/array_trie"
)

// Re[trie]ve is a effiecient data structure
// for storing and retieving words

func main() {
	myTrie := InitArrayTrie()
	
	myTrie.Insert("aragon")
	myTrie.Insert("a")

	fmt.Println(myTrie.Search("aragon"))
	fmt.Println(myTrie.Search("arag"))
	fmt.Println(myTrie.Search("a"))
}
