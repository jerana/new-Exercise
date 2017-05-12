package main

import "fmt"
import "strings"

func main() {
	path := "/a/./b///../c/../././../d/..//../e/./f/./g/././//.//h///././/..///"
	fmt.Println("find PAth:", simplifyPath(path))
}

/*
Given an absolute path for a file (Unix-style), simplify it.

For example,
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"
click to show corner cases.

Corner Cases:
Did you consider the case where path = "/../"?
In this case, you should return "/".
Another corner case is the path might contain multiple slashes '/' together, such as "/home//foo/".
In this case, you should ignore redundant slashes and return "/home/foo".

*/

func simplifyPath(path string) string {
	var list []string
	var subString []string
	var newPath string
	if len(path) > 0 && path[len(path)-1] == '/' {
		newPath = path[:len(path)-1]
	}
	subString = strings.SplitAfter(newPath, "/")
	var result string = ""
	for i, s := range subString {
		fmt.Println("num:", i, s)
		if (s == "./") || (s == "/") {
			fmt.Println("skip:", s)
		} else if s == "../" {
			l := len(list)
			fmt.Println("len:", l)
			//fmt.Println("del:",list[l])

			if l > 1 {
				list = list[:(l - 1)]
			}
			//a = a[:len(a)-1]
			fmt.Println(" after del node:", list)
		} else {
			list = append(list, s)
		}
		fmt.Println(" node:", list)
	}
	for _, r := range list {
		result += r
	}

	return result
}
