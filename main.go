package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Homework/Week3/NamikoToriyama/promode"
	"github.com/Homework/Week3/NamikoToriyama/trie"
)

var sc = bufio.NewScanner(os.Stdin)
var dictionary map[string]string
var dictArray []string
var filename = "dictionary/words.txt"
var trieTree *trie.Node

func main() {
	var err error
	dictionary, dictArray, err = readLine(filename, '\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	trieTree = trie.CreateTrie(dictArray)

	var s string
	fmt.Println("Hello!!")
	for {
		fmt.Println("If you select [pro], dictionary changes pro mode.")
		fmt.Println("~~SELECT MENU~~")
		fmt.Printf(" Search\twords\t[s]\n Add\twords\t[a]\n Delete\twords\t[d]\n Update\twords\t[u]\n Save\t\t[sv]\n Quit\t\t[q]\n> ")
		if sc.Scan() {
			s = sc.Text()
		}
		switch s {
		case "pro":
			promode.ProMode()
		case "search", "s":
			GetWord()
		case "add", "a":
			AddWords()
		case "delete", "d":
			DeleteWord()
		case "update", "u":
			UpdateWord()
		case "save", "sv":
			SaveWords()
		case "quit", "q":
			fmt.Println("Bye:)")
			return
		}
	}

}

func readLine(filename string, delim byte) (map[string]string, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	dict := map[string]string{}
	for {
		line, err := reader.ReadString(delim)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		slice := strings.Split(line, "\t")
		c := ""
		for i, s := range slice { // tabが複数あった場合
			if i == 0 {
				continue
			}
			s = strings.TrimRight(s, "\n")
			c += s
		}
		v, exist := dict[slice[0]]
		// 単一の単語に複数の意味がある場合
		if exist {
			c = v + " / " + c
		}
		// 単一の意味に対して複数の単語がある場合
		if strings.Contains(slice[0], ",") {
			key := strings.Split(slice[0], ",")
			if len(key) == 0 {
				key = strings.Split(slice[0], ", ")
			}
			for _, k := range key {
				dict[k] = c
				dictArray = append(dictArray, k)
			}
		} else {
			dict[slice[0]] = c
			dictArray = append(dictArray, slice[0])
		}
	}
	return dict, dictArray, nil
}
