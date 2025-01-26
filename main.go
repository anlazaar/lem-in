package main

import (
	"fmt"
	"lem-in/Algorithms"
	"lem-in/GlobVar"
	"lem-in/Helpers"
	"lem-in/Utils"
	"os"
	"sort"
)	



func main() {
	args := os.Args[1:]
    if len(args) != 1 {
        return
    }

    dataBytes, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println("ERROR: invalid data format;", err)
		return
	}

	err = Utils.ParsingData(string(dataBytes))


	GlobVar.OriginalRooms = Helpers.CopyRoomsMap(GlobVar.Rooms)
	if err != nil {
		fmt.Println("ERROR: invalid data format;", err)
		return
	}

	Algorithms.FindValidPaths()
	GlobVar.AllValidPaths = append(GlobVar.AllValidPaths, GlobVar.ValidPaths)

	// Sorting strings by length
	sort.Slice(GlobVar.ValidPaths, func(i, j int) bool {
		return len(GlobVar.ValidPaths[i]) < len(GlobVar.ValidPaths[j])
	})

	shortestPathIndex := 0
	lessTurns, antsOrdred := Algorithms.OrderAnts(0)

	for i := 1; i < len(GlobVar.AllValidPaths); i++ {
		turns, ants := Algorithms.OrderAnts(i)
		if turns < lessTurns {
			antsOrdred = ants
			lessTurns = turns
			shortestPathIndex = i
		}
	}

	Utils.HandleExport(antsOrdred,lessTurns,shortestPathIndex, string(dataBytes))
	fmt.Println()
}