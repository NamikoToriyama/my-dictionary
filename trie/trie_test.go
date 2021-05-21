package trie

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/deckarep/golang-set"
)

func createDict() []string {
	file, err := os.Open("../dictionary/test.txt")
	if err != nil {
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var dictArray []string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}
		slice := strings.Split(line, "\t")
		dictArray = append(dictArray, slice[0])

	}
	return dictArray
}

func TestSearchTrieTree(T *testing.T) {
	dict := createDict()
	tree := CreateTrie(dict)

	want := mapset.NewSetFromSlice([]interface{}{"decide", "decided"})
	var infc []interface{} = SearchTrieTree("deci", tree)
	got := mapset.NewSetFromSlice(infc)
	if want != got {
		t.Fatalf("want = %v, got = %v", want, got)
	}

	want = mapset.NewSetFromSlice([]interface{}{"decimal", "decimal fraction", "decimal point", "decimal system", "decimalise", "decimalize"})
	SearchTrieTree("decimal", tree)
	SearchTrieTree("d", tree)
	// t.Fatalf("want = %v, got = %v", want, got)
}
