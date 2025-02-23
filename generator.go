package main

func generateDungeon(numberOfFloors int) [][][]Cell {
	var dungeon [][][]Cell = make([][][]Cell, numberOfFloors)
	for i := 0; i < numberOfFloors; i++ {
		dungeon[i] = generateGrid(4, 6)
	}
	return dungeon
}

func generateGrid(sizeX, sizeY int) [][]Cell {
	var grid [][]Cell = make([][]Cell, sizeX)
	for i := 0; i < sizeX; i++ {
		grid[i] = make([]Cell, sizeY)
	}

	//temp
	/*grid = [][]Cell{
		{
			{Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 1 ", Filler: Filler{Empty: false}},
		},
		{
			{Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 1 ", Filler: Filler{Empty: false}},
		},
		{
			{Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 1 ", Filler: Filler{Empty: false}},
		},
		{
			{Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 0 ", Filler: Filler{Empty: true}}, {Symbol: " 1 ", Filler: Filler{Empty: false}},
		},
	}*/ //temp

	grid = [][]Cell{
		{
			{Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 1 ", Filler: Filler{Empty: false}},
		},
		{
			{Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: " 1 ", Filler: Filler{Empty: false}},
		},
		{
			{Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: " 1 ", Filler: Filler{Empty: false}}, {Symbol: " 1 ", Filler: Filler{Empty: false}},
		},
		{
			{Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: "   ", Filler: Filler{Empty: true}}, {Symbol: " 1 ", Filler: Filler{Empty: false}},
		},
	}

	return grid
}
