package main

import (
	"fmt"

	utils "lem_in/utils"
)

func main() {
	var (
		StartRoom  []string
		EndRoom    []string
		AntsNumber int
		GraphMap   = map[string][]string{}
	)
	StartRoom, EndRoom, AntsNumber, GraphMap = utils.GetFileData("data.txt")
	fmt.Println(StartRoom)
	fmt.Println(EndRoom)
	fmt.Println(AntsNumber)
	fmt.Println(DFSearch(StartRoom[0], EndRoom[0], []string{}, GraphMap, [][]string{}, map[string]bool{}))
}

func DFSearch(StartRoom, EndRoom string, Path []string, GraphMap map[string][]string, Result [][]string, Visited map[string]bool) [][]string {
	Visited[StartRoom] = true
	Path = append(Path, StartRoom)
	if StartRoom == EndRoom {
		PathCopy := make([]string, len(Path))
		copy(PathCopy, Path)
		return append(Result, PathCopy)
	}
	for _, RoomToGo := range GraphMap[StartRoom] {
		Result = DFSearch(RoomToGo, EndRoom, Path, GraphMap, Result, Visited)
		Path = Path[:len(Path)-1]
	}
	return Result
}
