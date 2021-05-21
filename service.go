package main

import "fmt"

// AddWords ... handles add service.
func AddWords() {
	var s string
	fmt.Printf("どの単語を登録しますか? > ")
	if sc.Scan() {
		s = sc.Text()
	}
	if s == "" {
		return
	}
	AddDictionary(s)
}

// GetWord ... handles get(search) service.
func GetWord() {
	var s string
	for {
		fmt.Printf("Input search Word > ")
		if sc.Scan() {
			s = sc.Text()
		}
		if s == "" {
			return
		}
		GetDictionary(s)
	}
}

// DeleteWord ... handles delete service.
func DeleteWord() {
	var s string
	fmt.Printf("Input delete word > ")
	if sc.Scan() {
		s = sc.Text()
	}
	_, exists := dictionary[s]
	if !exists {
		fmt.Println(s + ": is not registered")
		return
	}
	DeleteDictionary(s)

}

// UpdateWord ... handles update service.
func UpdateWord() {
	var s string
	fmt.Printf("Input update Word > ")
	if sc.Scan() {
		s = sc.Text()
	}
	if s == "" {
		return
	}
	k, exists := dictionary[s]
	if exists {
		fmt.Println("\n今までの意味は\n" + k + "\nとなっています。")
		UpdateDictionary(s)
		fmt.Println("意味が更新されました")
	} else {
		fmt.Println(s + ": is not registered")
		ConfirmRegister(s)
	}
}
