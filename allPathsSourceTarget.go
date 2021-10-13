package main

import "fmt"

//输入：graph = [[4, 3, 1], [3, 2, 4],[3], [4], []]
//输出：[[0, 4], [0, 3, 4], [0, 1, 3, 4], [0, 1, 2, 3, 4], [0, 1, 4]]

func findNextNode(graph [][]int, currentIndex int, target int, path []int, res *[][]int) {
	lenElement := len(graph[currentIndex])
	//graph := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
	for i := 0; i < lenElement; i++ {
		//当前节点包含最终节点
		if graph[currentIndex][i] == target {
			newPath := make([]int, len(path))
			copy(newPath, path)
			newPath = append(newPath, target)
			*res = append(*res, newPath)
			continue
		}
		newPath := append(path, graph[currentIndex][i])
		findNextNode(graph, graph[currentIndex][i], target, newPath, res)
	}
}
func allPathsSourceTarget(graph [][]int) [][]int {
	path := []int{0}
	target := len(graph)
	res := [][]int{}
	findNextNode(graph, 0, target-1, path, &res)
	return res
}
func main() {
	graph := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}
