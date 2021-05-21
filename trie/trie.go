package trie

// Node ... node for one alphabet.
type Node struct {
	value  string
	isWord bool
	word   string
	next   []*Node
}

// CreateTrie ... remake string array to trie tree.
func CreateTrie(dictionary []string) *Node {
	HEAD := &Node{"", false, "", []*Node{}}
	for _, word := range dictionary { // 単語ずつに分解
		InsertTrie(word, HEAD)
	}
	return HEAD
}

// InsertTrie ... insert word to trie tree.
func InsertTrie(word string, tree *Node) error {
	current := tree
	for _, c := range word { // charに分解
		np := &Node{string(c), false, current.word + string(c), []*Node{}}
		if e, n := isExist(string(c), current.next); e {
			np = n
		} else {
			current.next = append(current.next, np)
		}
		current = np // 次のノードに移動
	}
	current.isWord = true
	return nil
}

// SearchTrieTree ... search prefix word by using trie tree
func SearchTrieTree(word string, tree *Node) []string {
	current := tree
	for _, c := range word { // charに分解
		if e, n := isExist(string(c), current.next); e {
			current = n
		} else {
			return nil
		}
	}

	// BFSでその下を探すことで前方一致の単語を見つける
	prefixWords := bfs(current)
	printResult(word, prefixWords)

	return prefixWords
}
