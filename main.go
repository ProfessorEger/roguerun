package main

func main() {
	dungeon := generateDungeon(3)

	//dungeon := readDungeonFromFile("test.txt")
	for _, grid := range dungeon {
		printGrid(grid)
	}
}
