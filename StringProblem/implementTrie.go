package main

/*
Problem Statement:
Implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
Note:

You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.

*/
type Trie struct {
	found bool
	hMap  map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	var root Trie
	root.hMap = make(map[byte]*Trie, 0)
	return root
}

func allocTrieNode() *Trie {
	n := new(Trie)
	n.hMap = make(map[byte]*Trie, 0)
	return n
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	ok := false
	var key *Trie
	for _, w := range word {
		if key, ok = node.hMap[byte(w)]; !ok {
			n := allocTrieNode()
			node.hMap[byte(w)] = n
			node = n
		} else {
			node = key
		}
	}
	node.found = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.find(word)
	if node != nil && node.found {
		return true
	}
	return false
}
func (this *Trie) find(word string) *Trie {
	node := this
	for i := 0; i < len(word) && node != nil; i++ {
		node, _ = node.hMap[byte(word[i])]
	}
	return node
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.find(prefix)
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
