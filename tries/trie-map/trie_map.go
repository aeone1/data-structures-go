package trie_map

const AlphabetSize = 26

// Trie structure
// With an array
type Trie struct{
	root *Node
}

type Node struct {
	children map[rune]*Node
	isEnd bool
}

// Insert will take in a word and add it tot the trie
func (t *Trie) Insert(w string) {
	currentNode := t.root
	for _, char := range w {
		if child, ok := currentNode.children[char]; ok {
			currentNode = child
		} else {
			currentNode.children[char] = &Node{children: make(map[rune]*Node)}
			currentNode = currentNode.children[char]
		}
	}
	currentNode.isEnd = true
}

// Search will take in a word and return true if the word is included in the trie
func (t *Trie) Search(w string) bool {
	currentNode := t.root
	for _, char := range w {
		if child, ok := currentNode.children[char]; ok {
			currentNode = child
		} else {
			return false
		}
	}
	return currentNode.isEnd
}

// Initialize a trie
func InitTrie() *Trie {
	return &Trie{root: &Node{children: map[rune]*Node{}}}
}
