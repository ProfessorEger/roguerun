package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func printGrid(grid [][]Cell) {
	var strGrid = make([][]string, len(grid)*2+1)
	for i := 0; i < len(grid)*2+1; i++ {
		strGrid[i] = make([]string, len(grid[0])*2+1)
	}

	updateGrid(grid)
	fillStrGrid(strGrid, grid)
	addHorizotalLines(strGrid, grid)
	addVerticalLines(strGrid, grid)
	addCorners(strGrid)

	//clearConsole()
	printStrGrid(strGrid)
}

func updateGrid(grid [][]Cell) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !grid[i][j].Filler.Empty {
				grid[i][j].Symbol = grid[i][j].Filler.Symbol
			} else if grid[i][j].Creature.IsCreature {
				grid[i][j].Symbol = grid[i][j].Creature.Symbol
			} else if grid[i][j].Object.IsObject {
				grid[i][j].Symbol = grid[i][j].Object.Symbol
			} else {
				grid[i][j].Symbol = EMPTY_SYMBOL
			}
		}
	}
}

func fillStrGrid(strGrid [][]string, grid [][]Cell) {
	for i := 1; i < len(strGrid); i = i + 2 {
		for j := 1; j < len(strGrid[0]); j = j + 2 {
			strGrid[i][j] = grid[(i-1)/2][(j-1)/2].Symbol
		}
	}
}

func addHorizotalLines(strGrid [][]string, grid [][]Cell) {
	for i := 0; i < len(strGrid); i = i + 2 {
		for j := 1; j < len(strGrid[0]); j = j + 2 {
			if (((i/2)-1 >= 0 && !grid[(i/2)-1][(j-1)/2].Filler.Empty) && !(i/2 < len(grid) && !grid[i/2][(j-1)/2].Filler.Empty)) || ((i/2 < len(grid) && !grid[i/2][(j-1)/2].Filler.Empty) && !((i/2)-1 >= 0 && !grid[(i/2)-1][(j-1)/2].Filler.Empty)) {
				strGrid[i][j] = "───"
			} else {
				strGrid[i][j] = "   "
			}
		}
	}
}

func addVerticalLines(strGrid [][]string, grid [][]Cell) {
	for i := 1; i < len(strGrid); i = i + 2 {
		for j := 0; j < len(strGrid[0]); j = j + 2 {
			if (((j/2)-1 >= 0 && !grid[(i-1)/2][(j/2)-1].Filler.Empty) && !(j/2 < len(grid[0]) && !grid[(i-1)/2][j/2].Filler.Empty)) || ((j/2 < len(grid[0]) && !grid[(i-1)/2][j/2].Filler.Empty) && !((j/2)-1 >= 0 && !grid[(i-1)/2][(j/2)-1].Filler.Empty)) {
				strGrid[i][j] = "│"
			} else {
				strGrid[i][j] = " "
			}
		}
	}
}

func addCorners(strGrid [][]string) {
	for i := 0; i < len(strGrid); i = i + 2 {
		for j := 0; j < len(strGrid[0]); j = j + 2 {
			if i > 0 && j > 0 && i < len(strGrid)-1 && j < len(strGrid[0])-1 && checksLine(strGrid[i-1][j]) && checksLine(strGrid[i+1][j]) && checksLine(strGrid[i][j-1]) && checksLine(strGrid[i][j+1]) {
				strGrid[i][j] = "┼"
			} else if i > 0 && i < len(strGrid)-1 && checksLine(strGrid[i-1][j]) && checksLine(strGrid[i+1][j]) {
				strGrid[i][j] = "│"
			} else if j > 0 && j < len(strGrid[0])-1 && checksLine(strGrid[i][j-1]) && checksLine(strGrid[i][j+1]) {
				strGrid[i][j] = "─"
			} else if i < len(strGrid)-1 && j < len(strGrid[0])-1 && checksLine(strGrid[i+1][j]) && checksLine(strGrid[i][j+1]) {
				strGrid[i][j] = "┌"
			} else if i > 0 && j > 0 && checksLine(strGrid[i-1][j]) && checksLine(strGrid[i][j-1]) {
				strGrid[i][j] = "┘"
			} else if j > 0 && i < len(strGrid)-1 && checksLine(strGrid[i+1][j]) && checksLine(strGrid[i][j-1]) {
				strGrid[i][j] = "┐"
			} else if i > 0 && j < len(strGrid[0])-1 && checksLine(strGrid[i-1][j]) && checksLine(strGrid[i][j+1]) {
				strGrid[i][j] = "└"
			} else {
				strGrid[i][j] = " "
			}
		}
	}
}

func checksLine(str string) bool {
	return str == "│" || str == "───"
}

func printStrGrid(strGrid [][]string) {
	for _, row := range strGrid {
		for _, value := range row {
			fmt.Print(value)
		}
		fmt.Println()
	}
}

func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
