package lem_in

func DFSearch(StartRoom, EndRoom string, Path []string, GraphMap map[string][]string, Result [][]string) [][]string {
	Path = append(Path, StartRoom)
	if StartRoom == EndRoom {
		TemporaryPath := make([]string, len(Path))
		copy(TemporaryPath, Path)
		return append(Result, TemporaryPath)
	}
	for _, RoomToGo := range GraphMap[StartRoom] {
		if !Contian(Path, RoomToGo) {
			Result = DFSearch(RoomToGo, EndRoom, Path, GraphMap, Result)
		}
	}
	return Result
}

func Contian(Path []string, RoomToGo string) bool {
	for i := 0; i < len(Path); i++ {
		if Path[i] == RoomToGo {
			return true
		}
	}
	return false
}
