package practics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_Insert(t1 *testing.T) {
	trie := NewTrie()
	trie.insert("apple")
	assert.Equal(t1, true, trie.search("apple"))   // return True
	assert.Equal(t1, false, trie.search("app"))    // return False
	assert.Equal(t1, true, trie.startsWith("app")) // return True
	trie.insert("app")
	assert.Equal(t1, true, trie.search("app")) // return True
}
