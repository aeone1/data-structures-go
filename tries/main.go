package main

import (
	"fmt"
	arrayTrie "github.com/aeone1/data-structures-go/tries/trie-array"
	mapTrie "github.com/aeone1/data-structures-go/tries/trie-map"
)

// Re[trie]ve is a effiecient data structure
// for storing and retieving words

func main() {
	myArrayTrie := arrayTrie.InitTrie()
	
	myArrayTrie.Insert("aragon")
	myArrayTrie.Insert("a")

	fmt.Println(myArrayTrie.Search("aragon"))
	fmt.Println(myArrayTrie.Search("arag"))
	fmt.Println(myArrayTrie.Search("a"))

	myMapTrie := mapTrie.InitTrie()
	
	myMapTrie.Insert("aragon")
	myMapTrie.Insert("a")

	fmt.Println(myMapTrie.Search("aragon"))
	fmt.Println(myMapTrie.Search("arag"))
	fmt.Println(myMapTrie.Search("a"))
}
