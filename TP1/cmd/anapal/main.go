package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
)

func isPalindrome(word string) bool {
	for i := 0 ; i <= len(word)/2 ; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}

func Palindromes(words []string) (l []string) {
	for _, v := range words {
		if isPalindrome(v) {
			l = append(l, v)
		}
	} 
	return
}

func Footprint(s string) (footprint string) {
	copy_s := strings.Split(s,"")
	sort.Strings(copy_s)
	footprint = strings.Join(copy_s, "")
	return 
}

func Anagrams(words []string) (anagrams map[string][]string){
	anagrams = make(map[string][]string)
	for _, v := range words {
		anagrams[Footprint(v)] = append(anagrams[Footprint(v)], v)
	}
	return 
}

func DictFromFile(filename string) (dict []string) {
	file, err := os.ReadFile(filename) //on pourrait aussi utiliser open, mais necessite un buffer (plus précis mais moins rapide à mettre en place)
	if err != nil { //readfile ferme tout seul le fichier après son utilisation
		log.Fatal(err)
		return
	}
	copyf := make([]byte, len(file))
	copy(copyf, file)

	var temp_word_sl []string
	var temp_word string
	
	for _, v := range copyf {
		if v != 10 { //EOL rune number
			temp_word_sl = append(temp_word_sl, string(v))
		} else {
			temp_word = strings.Join(temp_word_sl, "") //join en spécifiant ce qu'il va y avoir entre chaque element de la slice !!
			dict = append(dict, temp_word)
			temp_word_sl = nil
		}
	}
	return
}

func maxSliceStr(sl []string) (max int) {
	for _, v := range sl {
		if len(v) > max {
			max = len(v)
		}
	}
	return
}

func isMaxPalindrome(dict []string) (res []string) {
	pal := Palindromes(dict)
	var palmax []string
	for _, v := range pal {
		if len(v) == maxSliceStr(pal) {
			palmax = append(palmax, v)
		}
	}
	return palmax
}

func anagramAgentsv1(dict []string) (res []string) { //avec double boucle for
	anagramap := Anagrams(dict)
	for _,v := range anagramap{ //i correspond à la clé, et v correspond à la valeur
		for _, w := range v {
			if w == "AGENTS" {
			return v
		}
		}
		
	}
	return nil
}

func anagramAgentsv2(dict []string) (res []string) { //avec slice.Contains
	anagramap := Anagrams(dict)
	for _,v := range anagramap{ //i correspond à la clé, et v correspond à la valeur
		if slices.Contains(v, "AGENTS") {
			return v
		}
	}
	return nil
}

func mostAnag(dict []string) (anagrams map[string][]string) {
	anagramap := Anagrams(dict)
	var max int
	anagrams = make(map[string][]string)
	for _, v := range anagramap {
		if len(v) > max {
			max = len(v)
		}
	}

	for i, v := range anagramap {
		if len(v) == max {
			anagrams[i] = v
		}
	}
	return anagrams
}

func palAnag(dict []string) (res map[string][]string){
	pal := Palindromes(dict)
	anags := Anagrams(pal)
	res = make(map[string][]string)
	for i,v := range anags {
		if len(v) >= 2{
			res[i] = v
		}
	}
	return
}


func main() {
	// dict := [...]string{"AGENT", "CHIEN", "COLOC", "ETANG", "ELLE", "GEANT", "NICHE", "RADAR"}
	// fmt.Print(isPalindrome("CHIEN"))
	// fmt.Println(dict)
	// fmt.Println(Palindromes(dict[:])) //passage en slice
	// fmt.Println(Footprint("COLOC"))
	// fmt.Println(Anagrams(dict[:]))
	new_dict := DictFromFile("c:/Users/utcpret/Downloads/dico-scrabble-fr.txt")
	// fmt.Print(isMaxPalindrome(new_dict))
	// fmt.Print(anagramAgentsv1(new_dict))
	// fmt.Print(anagramAgentsv2(new_dict))
	// fmt.Print(mostAnag(new_dict))
	fmt.Print(palAnag(new_dict))
}