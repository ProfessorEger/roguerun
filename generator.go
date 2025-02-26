package main

import (
	"fmt"
	"math/rand"
)

const (
	MIN_ROOM_SIZE = 3
	MIN_RECT_SIZE = 3
	MAX_GRID_SIZE = 29
	MIN_GRID_SIZE = 12
)

type Leaf struct {
	Size        [2]int
	Coordinates [2]int
	Room        Rect

	RightChild *Leaf
	LeftChild  *Leaf
}

func (l *Leaf) split() (isSmall bool) {
	width, height := l.Size[0], l.Size[1]
	var splitVertical bool

	if width > int(float64(height)*1.25) {
		splitVertical = true
	} else if height > int(float64(width)*1.25) {
		splitVertical = false
	} else {
		// Случайный выбор направления
		splitVertical = rand.Intn(2) == 0
	}

	maxSplit := 0
	if splitVertical {
		maxSplit = width - MIN_ROOM_SIZE
	} else {
		maxSplit = height - MIN_ROOM_SIZE
	}

	if maxSplit <= MIN_ROOM_SIZE {
		// Недостаточно места для разбиения
		return true
	}

	splitPos := rand.Intn(maxSplit-MIN_ROOM_SIZE) + MIN_ROOM_SIZE

	if splitVertical {
		l.LeftChild = &Leaf{
			Size:        [2]int{splitPos, height},
			Coordinates: l.Coordinates,
		}
		l.RightChild = &Leaf{
			Size:        [2]int{width - splitPos, height},
			Coordinates: [2]int{l.Coordinates[0] + splitPos, l.Coordinates[1]},
		}
	} else {
		l.LeftChild = &Leaf{
			Size:        [2]int{width, splitPos},
			Coordinates: l.Coordinates,
		}
		l.RightChild = &Leaf{
			Size:        [2]int{width, height - splitPos},
			Coordinates: [2]int{l.Coordinates[0], l.Coordinates[1] + splitPos},
		}
	}

	return false
}

func (l *Leaf) buildRect() {
	// Случайные размеры прямоугольника
	rectWidth := rand.Intn(l.Size[0]-MIN_RECT_SIZE+1) + MIN_RECT_SIZE
	rectHeight := rand.Intn(l.Size[1]-MIN_RECT_SIZE+1) + MIN_RECT_SIZE

	// Случайные координаты для размещения прямоугольника внутри Leaf
	rectX := rand.Intn(l.Size[0]-rectWidth+1) + l.Coordinates[0]
	rectY := rand.Intn(l.Size[1]-rectHeight+1) + l.Coordinates[1]

	// Создание прямоугольника и присвоение в поле Room
	l.Room = Rect{
		Size:        [2]int{rectWidth, rectHeight},
		Coordinates: [2]int{rectX, rectY},
	}

	fmt.Print(l.Room.Coordinates, l.Room.Size, l.Coordinates)
	fmt.Println()
}

/*
func (l *Leaf) addToSlice(leafs *[]Leaf) {
	if l == nil {
		return
	}

	*leafs = append(*leafs, *l) // Добавляем текущий лист
}*/

type Rect struct {
	Size        [2]int
	Coordinates [2]int
}

func generateDungeon(numberOfFloors int) [][][]Cell {
	var dungeon [][][]Cell = make([][][]Cell, numberOfFloors)
	for i := 0; i < numberOfFloors; i++ {
		dungeon[i] = generateGrid()
	}
	return dungeon
}

func generateGrid() [][]Cell {
	root := generateRootLeaf()
	fmt.Print(root.Size)
	fmt.Println()
	grid := createWallGrid(root.Size)
	smallLeafs := divideIntoSmall(root)
	generateRooms(smallLeafs)
	insertRooms(grid, smallLeafs)

	return grid
}

func createWallGrid(size [2]int) [][]Cell {
	var grid [][]Cell = make([][]Cell, size[0])
	for i := 0; i < size[0]; i++ {
		grid[i] = make([]Cell, size[1])

		for j := 0; j < size[1]; j++ {
			grid[i][j].Coordinates = [2]int{i, j}
			grid[i][j].Filler = fillerMap["1"]
			grid[i][j].Creature = creatureMap["0"]
			grid[i][j].Object = objectMap["0"]
		}
	}

	return grid
}

func divideIntoSmall(root *Leaf) []*Leaf {
	smallLeafs := []*Leaf{}
	readyToDivide := []*Leaf{}
	bufferReadyToDivide := []*Leaf{root}

	readyToDivide = append(readyToDivide, bufferReadyToDivide...)

	for {
		readyToDivide = readyToDivide[:0]
		readyToDivide = append(readyToDivide, bufferReadyToDivide...)
		bufferReadyToDivide = bufferReadyToDivide[:0]

		smallLeafsNumber := 0
		for i := 0; i < len(readyToDivide); i++ {
			if readyToDivide[i] != nil && !readyToDivide[i].split() {
				bufferReadyToDivide = append(bufferReadyToDivide, readyToDivide[i].RightChild, readyToDivide[i].LeftChild)
			} else {
				smallLeafs = append(smallLeafs, readyToDivide[i])
				smallLeafsNumber++
			}
			if smallLeafsNumber == len(readyToDivide) {
				return smallLeafs
			}
		}
	}
}

func generateRooms(leafs []*Leaf) {
	for i := 0; i < len(leafs); i++ {
		leafs[i].buildRect()
	}
}

func generateRootLeaf() *Leaf {
	width := rand.Intn(MAX_GRID_SIZE-MIN_GRID_SIZE+1) + MIN_GRID_SIZE
	height := rand.Intn(MAX_GRID_SIZE-MIN_GRID_SIZE+1) + MIN_GRID_SIZE

	return &Leaf{
		Size:        [2]int{width, height},
		Coordinates: [2]int{0, 0},
	}
}

func insertRooms(grid [][]Cell, leafs []*Leaf) {
	for _, leaf := range leafs {
		insertRoom(grid, leaf)
	}
}

func insertRoom(grid [][]Cell, leaf *Leaf) {
	minX := leaf.Room.Coordinates[0]
	minY := leaf.Room.Coordinates[1]
	size := leaf.Room.Size

	maxX := minX + size[0]
	maxY := minY + size[1]

	for i := minX; i < maxX; i++ {
		for j := minY; j < maxY; j++ {
			grid[i][j].Filler = fillerMap["0"]
		}
	}
}
