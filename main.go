package main

func main() {
	//var grid [][]Cell = generateGrid(4, 6)
	dungeon := readDungeonFromFile("test.txt")

	printGrid(dungeon[0])
	printGrid(dungeon[1])
	printGrid(dungeon[2])
}
