package trie_test

import (
	"testing"

	"github.com/claudiu/go-trie"
)

func TestAddAndFind(t *testing.T) {
	rootTrie := trie.NewTrie()
	rootTrie.Add("hello")
	rootTrie.Add("heart")

	if rootTrie.Find("hello") == nil {
		t.FailNow()
	}

	if rootTrie.Find("heart") == nil {
		t.FailNow()
	}

}

func TestExistance(t *testing.T) {
	rootTrie := trie.NewTrie()
	rootTrie.Add("heart")

	if rootTrie.Find("hearts") != nil {
		t.FailNow()
	}

	rootTrie.Add("hearts")

	if rootTrie.Find("hearts") == nil {
		t.FailNow()
	}

	if rootTrie.FindNode("he") == nil {
		t.FailNow()
	}
}

func TestInternational(t *testing.T) {
	rootTrie := trie.NewTrie()
	rootTrie.Add("многочисленное")

	if rootTrie.Find("многочисленное") == nil {
		t.FailNow()
	}

	rootTrie.Add("многочи")

	if rootTrie.Find("многочи") == nil {
		t.FailNow()
	}
}

func TestFindNode(t *testing.T) {
	rootTrie := trie.NewTrie()
	rootTrie.Add("tester")

	if rootTrie.FindNode("testing") != nil {
		t.Fail()
	}

	if rootTrie.FindNode("te") == nil {
		t.Fail()
	}

	if rootTrie.FindNode("test") == nil {
		t.Fail()
	}
}

func TestCountWords(t *testing.T) {
	rootTrie := trie.NewTrie()
	rootTrie.Add("ab")
	rootTrie.Add("abc")
	rootTrie.Add("abcd")
	rootTrie.Add("abcdefg")

	rootTrie.Add("zoidberg")
	rootTrie.Add("iceberg")

	if rootTrie.Count() != 6 {
		t.Fail()
	}
}

func TestMetaData(t *testing.T) {
	rootTrie := trie.NewTrie()
	rootTrie.Add("heart")

	node := rootTrie.Find("heart")

	node.Set("test", "alpha")
	item, exists := node.Get("test")

	if exists != true {
		t.FailNow()
	}

	if item.(string) != "alpha" {
		t.FailNow()
	}
}

func TestInexistantMetaData(t *testing.T) {
	rootTrie := trie.NewTrie()
	rootTrie.Add("heart")

	node := rootTrie.Find("heart")

	node.Set("test", "alpha")
	_, exists := node.Get("tests")

	if exists == true {
		t.FailNow()
	}
}
