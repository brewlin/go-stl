package test

import (
	"testing"

	"github.com/brewlin/go-stl/trie"
)

func TestTrieInsert(t *testing.T) {
	tr := trie.NewTrie()
	tr.Add("xiaodo")
	tr.Add("do")
}
func TestTrieContains(t *testing.T) {
	tr := trie.NewTrie()
	tr.Add("xiaodo")
	tr.Add("do")
	if !tr.Contains("xiaodo") {
		t.Error("查询失败")
	}
}

func TestTrieIsPre(t *testing.T) {
	tr := trie.NewTrie()
	tr.Add("xiaodo")
	tr.Add("do")
	if !tr.IsPrefix("xiao") {
		t.Error("前缀查询失败")
	}
}

func TestTrieSearch(t *testing.T) {
	tr := trie.NewTrie()
	tr.Add("xiaodo")
	tr.Add("do")
	if !tr.Search("xiaodo") {
		t.Error("搜索失败")
	}
	if !tr.Search("d.") {
		t.Error("搜索失败")
	}
}
