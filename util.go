package main

import (
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/Homework/Week3/NamikoToriyama/promode"
)

// ConfirmRegister ... confirm whether to register a word in the dictionary.
func ConfirmRegister(s string) {
	for {
		var t string
		fmt.Println("単語を登録しますか?[y or n] > ")
		if sc.Scan() {
			t = sc.Text()
		}
		if t == "y" {
			AddDictionary(s)
			break
		}
		if t == "n" {
			break
		}
	}
}

// SaveWords ... save new word you registered.
func SaveWords() error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	dictSlice := make([]string, len(dictionary))
	i := 0
	for k, v := range dictionary {
		line := k + "\t" + v + "\n"
		dictSlice[i] = line
		i++
	}
	sort.Slice(dictSlice, func(i, j int) bool { return dictSlice[i] < dictSlice[j] })

	for i := range dictSlice {
		_, err := file.WriteString(dictSlice[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// RecommendMeaning ... print meaning in weblio.
func RecommendMeaning(s string) {
	m, err := promode.DoScraping(s)
	if err != nil {
		return
	}
	fmt.Println("\nWeblioでは" + s + "の意味は、\n" + m + "\nと書かれています。")
}

func checkKey(index int, key string) bool {
	if dictArray[index] >= key {
		return true
	}
	return false
}

func printCandidateWords(key string, c []string) {
	if len(c) == 0 {
		fmt.Println()
		fmt.Println("/////////////////////////////////")
		fmt.Println(key + " と似ているワードはありませんでした。")
		fmt.Println("/////////////////////////////////")
		return
	}
	fmt.Println()
	fmt.Println("/////////////////////////////////")
	fmt.Println(key + " と似ているワード")
	fmt.Println("/////////////////////////////////")
	for _, w := range c {
		fmt.Println("- " + w)
	}
	fmt.Println()
}

// RecommendSimilarWords ... get similar 10 words by using binary search.
func RecommendSimilarWords(key string) {
	left := -1
	right := len(dictArray)
	var mid int
	for math.Abs(float64(right-left)) > 1 {
		mid = (left + right) / 2
		if checkKey(mid, key) {
			right = mid
		} else {
			left = mid
		}
	}

	// 前後のワードを5個ずつとってくる
	// ~10個までになるように候補のワードを絞る
	var candidateWords []string
	start := int(math.Max(float64(mid-5), 0))
	end := int(math.Min(float64(mid+5), float64(len(dictArray))))
	for i := start; i < mid; i++ {
		if dictArray[i][0] != key[0] {
			continue
		}
		candidateWords = append(candidateWords, dictArray[i])
	}

	for i := mid; i < end; i++ {
		if dictArray[i][0] != key[0] {
			continue
		}
		candidateWords = append(candidateWords, dictArray[i])
	}
	printCandidateWords(key, candidateWords)

}
