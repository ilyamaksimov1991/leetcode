package practics

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode, 0),
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
	return Trie{
		root: newTrieNode(),
	}
}

func (t *Trie) insert(word string) {
	trieNode := t.root
	for i := 0; i < len(word); i++ {
		if _, ok := trieNode.children[word[i]]; !ok {
			trieNode.children[word[i]] = newTrieNode()
		}
		trieNode = trieNode.children[word[i]]
	}

	trieNode.isEnd = true
}

func (t *Trie) search(word string) bool {
	trieNode := t.root
	for i := range word {
		if _, ok := trieNode.children[word[i]]; !ok {
			return false
		}
		trieNode = trieNode.children[word[i]]
	}

	return trieNode.isEnd
}

func (t *Trie) startsWith(prefix string) bool {
	trieNode := t.root
	for i := range prefix {
		if _, ok := trieNode.children[prefix[i]]; !ok {
			return false
		}
		trieNode = trieNode.children[prefix[i]]
	}

	return true
}
