package main

import "fmt"

//Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s) from beginWord to endWord, such that:

//Only one letter can be changed at a time
//Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
//For example,

//Given:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log","cog"]
//Return
// [
//    ["hit","hot","dot","dog","cog"],
//   ["hit","hot","lot","log","cog"]
// ]
//Note:
//Return an empty list if there is no such transformation sequence.
//All words have the same length.
//All words contain only lowercase alphabetic characters.
//You may assume no duplicates in the word list.
//You may assume beginWord and endWord are non-empty and are not the same.

func main() {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord := "hit"
	endWord := "cog"
	fmt.Println("result output:", findLadders(beginWord, endWord, wordList))
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	result := make([][]string, 0)
	unvisited := make(map[string]bool, 0)
	for _, w := range wordList {
		unvisited[w] = true
	}
	visited := make(map[string]bool, 0)
	var Queue []*wordNode
	newWord := new(wordNode)
	newWord.alloc(beginWord, 1, nil)

	Queue = append(Queue, newWord)
	minPath := 0
	prelabel := 0
	curlabel := 0
	for len(Queue) > 0 { //Till Queue is not empty

		word := Queue[0]
		Queue = Queue[1:]

		curlabel = word.numSteps
		if word.Node == endWord { //Reach Target Node
			if minPath == 0 {
				minPath = word.numSteps // Set minPath reached sofar
			}
			if word.numSteps == minPath && minPath != 0 {
				list := make([]string, 0)
				for word != nil {
					list = append(list, word.Node)
					word = word.preNode
				}
				result = append(result, reverseSlice(list))
				continue
			}

		}
		if prelabel < curlabel {
			for v, _ := range visited {
				delete(unvisited, string(v))
			}
		}
		prelabel = curlabel
		sWord := []byte(word.Node)
		for i := 0; i < len(word.Node); i++ {
			tmp := (sWord[i])
			for j := 0; j < 26; j++ {
				(sWord)[i] = byte('a' + j)
				lkupWord := (string)(sWord)
				if _, ok := unvisited[lkupWord]; ok {
					nWord := new(wordNode)
					nWord.alloc(lkupWord, word.numSteps+1, word)
					Queue = append(Queue, nWord)
					visited[lkupWord] = true
				}
			}
			[]byte(sWord)[i] = tmp
		}
	}
	return result
}

type wordNode struct {
	Node     string
	numSteps int
	preNode  *wordNode
}

func (w *wordNode) alloc(s string, step int, pre *wordNode) {
	w.Node = s
	w.numSteps = step
	w.preNode = pre
}

func reverseSlice(list []string) []string {
	last := len(list) - 1
	for i := 0; i < len(list)/2; i++ {
		list[i], list[last-i] = list[last-i], list[i]
	}
	return list
}
