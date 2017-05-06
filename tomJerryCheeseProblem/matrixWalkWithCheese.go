package main

import "fmt"

/*
There is a maze of size n*n. Tom is sitting at (0,0). Jerry is sitting in another cell (the position of Jerry is input).
Then there are k pieces of cheese placed in k different cells (k <= 10).
Some cells are blocked while some are not.
Tom can move to 4 cells at any point of time (left, right, up, down one position).
Tom has to collect all the pieces of cheese and then reach to Jerry’s cell.
You need to print the minimum no. of steps required to do so.
*/

/*Algorithm :
Step1 : First find all cheese location
Step2 : Find Minimum path  to reach one of above find location  from Tom origin and make it next origin position for next
recursive till you reach last find location
Step3 : find minimum path from last find location to Jerry position
*/

const maxInt32 = 1<<31 - 1

type points struct {
	x, y int
}

var visited [][]int

func main() {
	arr := [][]int{{0, 0, 0, 1, 2},
		{0, 1, 0, 1, 0},
		{1, 2, 0, 0, 0},
		{2, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
	}

	visited = make([][]int, 5)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, 5)
	}
	//Jerry Location :
	x, y := 4, 2
	fmt.Println("ShortestPath :", findMinPathToReachJerry(arr, x, y))
}
func visited_matrix_clean(visited [][]int) {
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[0]); j++ {
			visited[i][j] = 0
		}
	}
}

func findMinPathToReachJerry(matrix [][]int, x, y int) int {
	m := len(matrix)
	n := len(matrix[0])
	var cheeseLoc = make([]points, 0)
	var loc points
	totalMinPath := 0
	//Step1: Record all Cheese coordinates into list
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 2 {
				loc.x = i
				loc.y = j
				cheeseLoc = append(cheeseLoc, loc)
			}

		}
	}
	fmt.Println("Cheese List:", cheeseLoc)
	//Step2 : Get shortest combination from all points
	//Tom init position
	start_x := 0
	start_y := 0
	for len(cheeseLoc) > 1 { //Walk all over coordinate and find minimum path from origin to one of remaining coordinate

		minPath := maxInt32 //Init Value of MiniPath from origin
		nextOrigin := points{0, 0}
		rmIndex := 0 //Remove CheeseLocation index which found to nearest to define origin point
		for i := 0; i < len(cheeseLoc); i++ {
			loc := cheeseLoc[i]
			visited_matrix_clean(visited)
			path := findMinPath(matrix, start_x, start_y, loc.x, loc.y, visited)
			if path < minPath {
				nextOrigin.x = loc.x
				nextOrigin.y = loc.y
				rmIndex = i
				minPath = path
			}
		}
		//At this stage, Got nearest cheese location from define origin points
		//Update new origin points and find nearest cheese from this new origin
		start_x = nextOrigin.x
		start_y = nextOrigin.y
		cheeseLoc = append(cheeseLoc[:rmIndex], cheeseLoc[rmIndex+1:]...) //skip rmIndex
		totalMinPath += minPath
	}
	//At this stage , All cheese has been collected and their minMum path has been calculated
	//Now find minMumPath to reach Jerry from last Location of cheese
	visited_matrix_clean(visited)
	totalMinPath += findMinPath(matrix, start_x, start_y, cheeseLoc[0].x, cheeseLoc[0].y, visited)
	return totalMinPath
}
func findMinPath(matrix [][]int, x, y, target_x, target_y int, visited [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	if target_x == x && target_y == y {
		return 0
	}
	if x < 0 || x > m-1 || y < 0 || y > n-1 || matrix[x][y] == 1 || visited[x][y] == 1 {
		return maxInt32
	}
	visited[x][y] = 1 //mark as visited
	minPath := maxInt32
	var kShift = []points{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i := 0; i < len(kShift); i++ { //DFS walk in all four directions
		next_x := x + kShift[i].x
		next_y := y + kShift[i].y
		path := 1 + findMinPath(matrix, next_x, next_y, target_x, target_y, visited)
		if path < minPath {
			minPath = path
		}
	}
	return minPath
}
