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
	var result string
	var list []string
	if len(path) > 0 || path[0] == '/' {
		list = append(list, string(path[0]))
	}
	subDir := strings.Split(path, "/")
	for _, sDir := range subDir {
		if sDir == ".." {
			if len(list) == 0 || list[len(list)-1] == ".." {
				list = append(list, sDir)
			} else {
				if list[len(list)-1] == "/" {
					fmt.Println("Invalid Path passed")
				} else {
					list = append(list, sDir)
				}

			}
		} else if s != "." && s != " " {
			list = append(list, sDir)
		}

	}
	for i, r := range list {
		if i == 0 {
			result = r
		} else if i == 1 && result == "/" {
			result += r
		} else {
			result += "/" + r
		}
	}
}
