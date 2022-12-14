package main

import "fmt"


func uniqueMorseRepresentations(words []string) int {
	code := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	wordCodeMap :=make(map[string]int)
	for _, word := range words {
		wordCode := ""
		for _,i :=range word{
			wordCode += code[i-97]
		}
		wordCodeMap[wordCode] =1
	}
	return len(wordCodeMap)
}
func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))

}

