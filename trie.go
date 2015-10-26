package trie

import "sync"

// Trie Holder
type Trie struct {
	letter   rune
	mutex    sync.RWMutex
	children []*Trie
	meta     map[string]interface{}
	isLeaf   bool
	//parent   *Trie
}

// NewTrie creates a new Trie and initialize it with default options
// May be used to create a new root trie.
func NewTrie() *Trie {
	newTrie := &Trie{}
	newTrie.children = []*Trie{}
	newTrie.meta = make(map[string]interface{})

	return newTrie
}

func (tr *Trie) hasChild(a rune) (bool, *Trie) {
	for _, child := range tr.children {
		if child.letter == a {
			return true, child
		}
	}

	return false, nil
}

func (tr *Trie) addChild(a rune) *Trie {
	tr.mutex.Lock()

	nw := NewTrie()
	nw.letter = a
	//nw.parent = tr
	tr.children = append(tr.children, nw)

	tr.mutex.Unlock()

	return nw
}

// Count returns the number of words in a trie.
func (tr *Trie) Count() int {
	count := 0
	for _, child := range tr.children {
		if child.isLeaf == true {
			count++
		}

		count += child.Count()
	}

	return count
}

// Add a word to a trie
func (tr *Trie) Add(word string) *Trie {
	letters, node, i := []rune(word), tr, 0
	n := len(letters)

	for i < n {
		if exists, value := node.hasChild(letters[i]); exists {
			node = value
		} else {
			node = node.addChild(letters[i])
		}

		i++

		if i == n {
			tr.mutex.Lock()
			node.isLeaf = true
			tr.mutex.Unlock()
		}
	}

	return node
}

// Remove a word from a trie.
// This does not free memory, it just marks the node.
func (tr *Trie) Remove(word string) {
	a := tr.Find(word)

	if a != nil {
		tr.mutex.Lock()
		a.isLeaf = false
		tr.mutex.Unlock()
	}
}

// FindNode returns the node whether it is a word or not.
func (tr *Trie) FindNode(word string) *Trie {
	letters, node, i := []rune(word), tr, 0
	n := len(letters)

	for i < n {
		if exists, value := node.hasChild(letters[i]); exists {
			node = value
		} else {
			return nil
		}

		i++
	}

	return node
}

// Find returns the node pointing to a word
func (tr *Trie) Find(word string) *Trie {
	node := tr.FindNode(word)

	if node == nil {
		return nil
	}

	if node.isLeaf != true {
		return nil
	}

	return node
}

// Get metadata belonging to a node.
func (tr *Trie) Get(key string) (interface{}, bool) {
	if tr == nil {
		return nil, false
	}

	if _, ok := tr.meta[key]; ok {
		return tr.meta[key], true
	}

	return nil, false
}

// Set metadata for a word
func (tr *Trie) Set(key string, value interface{}) {
	if tr == nil {
		return
	}

	tr.mutex.Lock()
	tr.meta[key] = value
	tr.mutex.Unlock()
}
