[![Build Status](https://travis-ci.org/Claudiu/Trie.svg)](https://travis-ci.org/Claudiu/Trie)
[![GoDoc](https://camo.githubusercontent.com/bfdd3541106bf567a1c4339af2cbf33fc60257e2/68747470733a2f2f662e636c6f75642e6769746875622e636f6d2f6173736574732f343536362f313133353630352f61623439323939302d316331392d313165332d383633622d6466646337653635313766312e706e67)](https://godoc.org/github.com/Claudiu/Trie)
# Trie / Radix Tree / Digital Tree


In computer science, a trie, also called digital tree and sometimes radix tree or prefix tree (as they can be searched by prefixes), is an ordered tree data structure that is used to store a dynamic set or associative array where the keys are usually strings.

Read More on:
https://en.wikipedia.org/wiki/Trie

## Usage

You start by defining the root of the Trie. Using ```trie.NewTrie()```
```go
package main

import "github.com/claudiu/trie"

func main() {
    rootTrie := trie.NewTrie()
    rootTrie.Add("heart")

    node := rootTrie.Find("heart")

    // do stuff with the node here
}
```

You can now return the trie node using:
```go
rootTrie.Find("heart")
```

If the node is not found ```rootTrie.Find("heart")``` will return _nil_


### Word Count
You can return the word count in a trie using:
```go
rootTrie.Count()
```

### Metadata
Each Trie node can hold metadata. Here's a small demo:

```go
package main

import (
    "github.com/claudiu/trie"
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
