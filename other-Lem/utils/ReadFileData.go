package lem_in

import (
	"os"
	"strconv"
	"strings"
)

func GetFileData(FileName string) (StartRoom []string, EndRoom []string, AntsNumber int, GraphMap map[string][]string) {
	GraphMap = map[string][]string{}
	Data, err := os.ReadFile(FileName)
	if err != nil {
		return
	}
	DataLines := strings.Split(string(Data), "\n")
	DateLenght := len(DataLines)
	for i := 0; i < DateLenght; i++ {
		if i == 0 {
			AntsNumber, err = strconv.Atoi(DataLines[i])
			if err != nil {
				return
			}
		} else if DataLines[i] == "##start" && i != DateLenght-1 && len(strings.Split(DataLines[i+1], " ")) == 3 && !strings.HasPrefix(DataLines[i+1], "L") {
			StartRoom = strings.Split(DataLines[i+1], " ")
		} else if DataLines[i] == "##end" && i != DateLenght-1 && len(strings.Split(DataLines[i+1], " ")) == 3 && !strings.HasPrefix(DataLines[i+1], "L") {
			EndRoom = strings.Split(DataLines[i+1], " ")
		} else if len(strings.Split(DataLines[i], "-")) == 2 {
			Tunnel := strings.Split(DataLines[i], "-")
			if strings.HasPrefix(Tunnel[0], "L") || strings.HasPrefix(Tunnel[1], "L") {
				continue
			}
			GraphMap[Tunnel[0]] = append(GraphMap[Tunnel[0]], Tunnel[1])
			GraphMap[Tunnel[1]] = append(GraphMap[Tunnel[1]], Tunnel[0])
		} else {
			continue
		}
	}
	return
}
