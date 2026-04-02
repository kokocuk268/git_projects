package main

import (
	"fmt"
	"sort"
	"strings"
)

func AnalyzeText(text string) {
	//! 1. Количество слов в тексте (считаем, что все слова в тексте разделяются только пробелами, точками, запятыми, восклицательным или вопросительным знаками)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' ' || r == '.' || r == ',' || r == '!' || r == '?'
	})
	fmt.Println("Количество слов:", len(words))

	//! 2. Количество уникальных слов (регистр не влияет на уникальность слова, то есть Привет и привет считаются одним словом)
	wordCount := make(map[string]int)
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		wordCount[lowerWord]++
	}
	fmt.Println("Количество уникальных слов:", len(wordCount))

	//! 3. Слово, которое встретилось в тексте больше всего раз (без учёта регистра)
	var maxOfWord string
	maxOfWordInt := 0
	for key, value := range wordCount {
		if value > maxOfWordInt {
			maxOfWordInt = value
			maxOfWord = key
		}
	}
	fmt.Printf("Самое часто встречающееся слово: %q (встречается %d раз)\n", maxOfWord, maxOfWordInt)
	//!4
	topWords := getTopWords(wordCount, 5)
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for _, word := range topWords {
		count := wordCount[word]
		fmt.Printf("%q: %d раз\n", word, count)
	}
}

func getTopWords(wordMap map[string]int, n int) []string {
	//! 4. Топ-5 часто встречающихся слов, по убыванию.
	newSlice := make([]string, 0, len(wordMap))

	for key := range wordMap {
		newSlice = append(newSlice, key)
	}

	sort.Slice(newSlice, func(i, j int) bool {
		return wordMap[newSlice[i]] > wordMap[newSlice[j]]
	})

	if len(newSlice) < n {
		return newSlice
	}

	return newSlice[:n]
}