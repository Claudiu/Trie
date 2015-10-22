# Trie / Radix Tree / Digital Tree


In computer science, a trie, also called digital tree and sometimes radix tree or prefix tree (as they can be searched by prefixes), is an ordered tree data structure that is used to store a dynamic set or associative array where the keys are usually strings.

Read More on:
https://en.wikipedia.org/wiki/Trie

## Usage

You start by defining the root of the Trie. Using ```trie.NewTrie()```
```
package main

import "github.com/claudiu/go-trie"

func main() {
    rootTrie := trie.NewTrie()
    rootTrie.Add("heart")

    node := rootTrie.Find("heart")

    // do stuff with the node here
}
```

You can now return the trie node using:
```
rootTrie.Find("heart")
```

If the node is not found ```rootTrie.Find("heart")``` will return _nil_

### Metadata
Each Trie node can hold metadata. Here's a small demo:

```
package main

import (
    "github.com/claudiu/go-trie"
    "fmt"
)

func main() {
    rootTrie := trie.NewTrie()
    rootTrie.Add("heart")

    node := rootTrie.Find("heart")

    node.Set("test", "alpha")

    item, found := node.Get("test")

    if found {
        fmt.Println(item.(string))
    }
}
```
