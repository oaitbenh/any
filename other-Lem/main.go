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
	ShortestWhitoutSort := GetShortestPaths(AllPaths)
	sort.Slice(AllPaths, func(i, j int) bool {
		return len(AllPaths[i]) < len(AllPaths[j])
	})
	ShortestPaths := GetShortestPaths(AllPaths)
	if len(ShortestWhitoutSort) > len(ShortestPaths) {
		ShortestPaths = ShortestWhitoutSort
		sort.Slice(ShortestPaths, func(i, j int) bool {
			return len(ShortestPaths[i]) < len(ShortestPaths[j])
		})
	}
	fmt.Println(ShortestPaths)
	AntAndPath := make([][]int, len(ShortestPaths))
	LenPaths := len(ShortestPaths)
	i := 1
	for i <= AntsNumber {
		for j := 0; j < LenPaths; j++ {
			if j == LenPaths-1 || len(ShortestPaths[j])+len(AntAndPath[j]) <= len(ShortestPaths[j+1])+len(AntAndPath[j+1]) {
				AntAndPath[j] = append(AntAndPath[j], i)
				i++
				if i > AntsNumber {
					break
				}
			}
		}
	}
	fmt.Println(AntAndPath)
}

func GetShortestPaths(AllPaths [][]string) (ShortestPaths [][]string) {
	PathsLenght := len(AllPaths)
	for i := 0; i < PathsLenght; i++ {
		AppendIt := true
		for j := 0; j < len(ShortestPaths); j++ {
			if MatchAnyRoom(AllPaths[i], ShortestPaths[j]) {
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
