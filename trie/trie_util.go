package trie

import (
	"container/list"
	"fmt"
)

// valueがNode.nextに存在しているかを確認する
func isExist(v string, cn []*Node) (bool, *Node) {
	for _, n := range cn {
		if n.value == v {
			return true, n
		}
	}
	return false, nil
}

func bfs(current *Node) []string {
	queue := list.New()
	queue.PushBack(current.next)
	var ans []string

	for queue.Len() > 0 {
		qnode := queue.Front()
		for _, n := range qnode.Value.([]*Node) {
			if n.isWord {
				ans = append(ans, n.word)
				if len(ans) > 9 {
					return ans
				}
			}
			if len(n.next) > 0 {
				queue.PushBack(n.next)
			}
		}
		queue.Remove(qnode)
	}
	return ans
}

func printResult(word string, prefixWords []string) {
	fmt.Println()
	if len(prefixWords) == 0 {
		fmt.Println("/////////////////////////////////")
		fmt.Println(word + "と前方一致するワードはありませんでした。")
		fmt.Println("/////////////////////////////////")
		fmt.Println()
		return
	}
	fmt.Println("/////////////////////////////////")
	fmt.Println(word + "と前方一致するワード")
	fmt.Println("/////////////////////////////////")
	for _, a := range prefixWords {
		fmt.Println("- " + a)
	}
	fmt.Println()
}
