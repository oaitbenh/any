package main

import (
	"fmt"
	"sort"

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
	AllPaths := utils.DFSearch(StartRoom[0], EndRoom[0], []string{}, GraphMap, [][]string{})
	sort.Slice(AllPaths, func(i, j int) bool {
		return len(AllPaths[i]) < len(AllPaths[j])
	})
	fmt.Println(AllPaths)
	ShortestPaths := GetShortestPaths(AllPaths)
	sort.Slice(ShortestPaths, func(i, j int) bool {
		return len(AllPaths[i]) < len(AllPaths[j])
	})
	fmt.Println(ShortestPaths)
}

func GetShortestPaths(AllPaths [][]string) (ShortestPaths [][]string) {
	for i := len(AllPaths) - 1; i >= 0; i-- {
		AppendIt := true
		for j := i - 1; j >= 0; j-- {
			if MatchAnyRoom(AllPaths[i], AllPaths[j]) {
				AppendIt = false
				break
			}
		}
		if AppendIt {
			ShortestPaths = append(ShortestPaths, AllPaths[i])
		}
	}
	return
}

func MatchAnyRoom(RoomsOne, RoomsTwo []string) bool {
	for i := 1; i < len(RoomsOne)-1; i++ {
		for j := 1; j < len(RoomsTwo)-1; j++ {
			if RoomsOne[i] == RoomsTwo[j] && i != 0 && i != len(RoomsOne)-1 {
				return true
			}
		}
	}
	return false
}
