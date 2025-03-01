package main

import (
	"fmt"
	"roguerun/models"
	"roguerun/service/configreader"
	"roguerun/service/generator"
	"roguerun/service/user_interaction"
)

func main() {
	err := configreader.ApplyConfig("config.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	dungeon := generator.GenerateDungeon(models.NUMBER_OF_FLOORS)

	//dungeon = dungeon_reader.ReadDungeonFromFile("test.txt")
	for _, grid := range dungeon {
		user_interaction.PrintGrid(grid)
	}
}
