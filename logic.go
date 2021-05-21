package main

import (
	"fmt"

	"github.com/Homework/Week3/NamikoToriyama/trie"
)

// AddDictionary ... add a word to the dictionary.
func AddDictionary(s string) {
	_, exists := dictionary[s]

	if exists {
		fmt.Println("Already registered")
	} else {
		RecommendMeaning(s)
		fmt.Printf("単語の意味を教えてください([q]でやめることが出来ます) > ")
		if sc.Scan() {
			m := sc.Text()
			if m == "q" {
				return
			}
			dictionary[s] = m
			dictArray = append(dictArray, s)
			trie.InsertTrie(s, trieTree)
		}
	}
}

// GetDictionary ... get a word from the dictionary.
func GetDictionary(s string) {
	k, exists := dictionary[s]
	if exists {
		fmt.Println(k)
	} else {
		fmt.Println(s + ": is not registered")
		trie.SearchTrieTree(s, trieTree)
		RecommendSimilarWords(s)
		ConfirmRegister(s)
	}
}

// DeleteDictionary ... delete a word from the dictionary.
func DeleteDictionary(s string) {
	for {
		var t string
		fmt.Println(s + ": 本当にこの単語を消しますか?[y or n] > ")
		if sc.Scan() {
			t = sc.Text()
		}
		if t == "y" {
			delete(dictionary, s)
			break
		}
		if t == "n" {
			break
		}
	}
}

// UpdateDictionary ... update a word in the dictionary.
func UpdateDictionary(s string) {
	RecommendMeaning(s)
	fmt.Printf("Updateしたい単語の意味を教えてください([q]でやめることが出来ます) > ")
	if sc.Scan() {
		m := sc.Text()
		if m == "q" {
			return
		}
		dictionary[s] = m
	}
}
