package array_trie

const AlphabetSize = 26

// Trie structure
// With an array
type Trie struct{
	root *Node
}

type Node struct {
	children [AlphabetSize]*Node
	isEnd bool
}

// Insert will take in a word and add it tot the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and return true if the word is included in the trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

// Initialize a trie
func InitArrayTrie() *Trie {
	return &Trie{root: &Node{}}
}
