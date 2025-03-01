package generator

import (
	"math/rand"
	"roguerun/models"
)

type leaf struct {
	Size        [2]int
	Coordinates [2]int
	Room        rect

	RightChild *leaf
	LeftChild  *leaf
}

type rect struct {
	Size        [2]int
	Coordinates [2]int
}

func (l *leaf) split() (isSmall bool) {
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
		maxSplit = width - models.MIN_LEAF_SIZE
	} else {
		maxSplit = height - models.MIN_LEAF_SIZE
	}

	if maxSplit <= models.MIN_LEAF_SIZE {
		// Недостаточно места для разбиения
		return true
	}

	splitPos := rand.Intn(maxSplit-models.MIN_LEAF_SIZE) + models.MIN_LEAF_SIZE

	if splitVertical {
		l.LeftChild = &leaf{
			Size:        [2]int{splitPos, height},
			Coordinates: l.Coordinates,
		}
		l.RightChild = &leaf{
			Size:        [2]int{width - splitPos, height},
			Coordinates: [2]int{l.Coordinates[0] + splitPos, l.Coordinates[1]},
		}
	} else {
		l.LeftChild = &leaf{
			Size:        [2]int{width, splitPos},
			Coordinates: l.Coordinates,
		}
		l.RightChild = &leaf{
			Size:        [2]int{width, height - splitPos},
			Coordinates: [2]int{l.Coordinates[0], l.Coordinates[1] + splitPos},
		}
	}

	return false
}

func (l *leaf) buildRect() {
	// Случайные размеры прямоугольника
	rectWidth := rand.Intn(l.Size[0]-models.MIN_RECT_SIZE+1) + models.MIN_RECT_SIZE
	rectHeight := rand.Intn(l.Size[1]-models.MIN_RECT_SIZE+1) + models.MIN_RECT_SIZE

	// Случайные координаты для размещения прямоугольника внутри Leaf
	rectX := rand.Intn(l.Size[0]-rectWidth+1) + l.Coordinates[0]
	rectY := rand.Intn(l.Size[1]-rectHeight+1) + l.Coordinates[1]

	// Создание прямоугольника и присвоение в поле Room
	l.Room = rect{
		Size:        [2]int{rectWidth, rectHeight},
		Coordinates: [2]int{rectX, rectY},
	}
}

func GenerateDungeon(numberOfFloors int) [][][]models.Cell {
	var dungeon [][][]models.Cell = make([][][]models.Cell, numberOfFloors)
	for i := 0; i < numberOfFloors; i++ {
		dungeon[i] = generateGrid()
	}
	return dungeon
}

func generateGrid() [][]models.Cell {
	root := generateRootLeaf()
	grid := createWallGrid(root.Size)
	smallLeafs := divideIntoSmall(root)
	generateRooms(smallLeafs)
	insertRooms(grid, smallLeafs)
	connectAllRooms(grid, smallLeafs)
	addBorders(grid)

	return grid
}

func createWallGrid(size [2]int) [][]models.Cell {
	var grid [][]models.Cell = make([][]models.Cell, size[0])
	for i := 0; i < size[0]; i++ {
		grid[i] = make([]models.Cell, size[1])

		for j := 0; j < size[1]; j++ {
			grid[i][j].Coordinates = [2]int{i, j}
			grid[i][j].Filler = models.FillerMap["1"]
			grid[i][j].Creature = models.CreatureMap["0"]
			grid[i][j].Object = models.ObjectMap["0"]
		}
	}

	return grid
}

func addBorders(grid [][]models.Cell) {
	for j := 0; j < len(grid[0]); j++ {
		grid[0][j].Filler = models.FillerMap["1"]
		grid[len(grid)-1][j].Filler = models.FillerMap["1"]
	}
	for i := 0; i < len(grid); i++ {
		grid[i][0].Filler = models.FillerMap["1"]
		grid[i][len(grid[0])-1].Filler = models.FillerMap["1"]
	}
}

func divideIntoSmall(root *leaf) []*leaf {
	smallLeafs := []*leaf{}
	readyToDivide := []*leaf{}
	bufferReadyToDivide := []*leaf{root}

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

func generateRooms(leafs []*leaf) {
	for i := 0; i < len(leafs); i++ {
		leafs[i].buildRect()
	}
}

func generateRootLeaf() *leaf {
	width := rand.Intn(models.MAX_GRID_SIZE-models.MIN_GRID_SIZE+1) + models.MIN_GRID_SIZE
	height := rand.Intn(models.MAX_GRID_SIZE-models.MIN_GRID_SIZE+1) + models.MIN_GRID_SIZE

	return &leaf{
		Size:        [2]int{width, height},
		Coordinates: [2]int{0, 0},
	}
}

func insertRooms(grid [][]models.Cell, leafs []*leaf) {
	for _, leaf := range leafs {
		insertRoom(grid, leaf)
	}
}

func insertRoom(grid [][]models.Cell, leaf *leaf) {
	minX := leaf.Room.Coordinates[0]
	minY := leaf.Room.Coordinates[1]
	size := leaf.Room.Size

	maxX := minX + size[0]
	maxY := minY + size[1]

	for i := minX; i < maxX; i++ {
		for j := minY; j < maxY; j++ {
			grid[i][j].Filler = models.FillerMap["0"]
		}
	}
}

// Функция для соединения всех комнат в подземелье
func connectAllRooms(grid [][]models.Cell, leafs []*leaf) {
	// Проверка на наличие как минимум двух комнат
	if len(leafs) < 2 {
		return
	}

	// Создаем граф соединенных комнат
	connected := make(map[*leaf]bool)
	connected[leafs[0]] = true

	// Соединяем все комнаты
	for len(connected) < len(leafs) {
		bestDistance := -1
		var roomA, roomB *leaf

		// Находим ближайшую пару комнат, одна из которых соединена, а другая нет
		for connectedLeaf := range connected {
			for _, otherLeaf := range leafs {
				if !connected[otherLeaf] {
					// Вычисляем расстояние между центрами комнат
					distance := calculateDistance(connectedLeaf, otherLeaf)

					if bestDistance == -1 || distance < bestDistance {
						bestDistance = distance
						roomA = connectedLeaf
						roomB = otherLeaf
					}
				}
			}
		}

		if roomA != nil && roomB != nil {
			// Соединяем найденную пару комнат коридором
			connectRoomsDirectly(grid, roomA, roomB)
			connected[roomB] = true
		}
	}
}

// Функция для вычисления расстояния между центрами комнат
func calculateDistance(leafA, leafB *leaf) int {
	centerAX := leafA.Room.Coordinates[0] + leafA.Room.Size[0]/2
	centerAY := leafA.Room.Coordinates[1] + leafA.Room.Size[1]/2
	centerBX := leafB.Room.Coordinates[0] + leafB.Room.Size[0]/2
	centerBY := leafB.Room.Coordinates[1] + leafB.Room.Size[1]/2

	// Манхэттенское расстояние
	return absInt(centerAX-centerBX) + absInt(centerAY-centerBY)
}

// Вспомогательная функция для вычисления абсолютного значения
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Функция для соединения двух комнат напрямую
func connectRoomsDirectly(grid [][]models.Cell, leafA, leafB *leaf) {
	// Получаем центры комнат
	startX := leafA.Room.Coordinates[0] + leafA.Room.Size[0]/2
	startY := leafA.Room.Coordinates[1] + leafA.Room.Size[1]/2
	endX := leafB.Room.Coordinates[0] + leafB.Room.Size[0]/2
	endY := leafB.Room.Coordinates[1] + leafB.Room.Size[1]/2

	// Определяем случайно, в каком порядке строить коридор:
	// сначала по горизонтали, потом по вертикали, или наоборот
	if rand.Intn(2) == 0 {
		createHorizontalCorridor(grid, startX, endX, startY)
		createVerticalCorridor(grid, startY, endY, endX)
	} else {
		createVerticalCorridor(grid, startY, endY, startX)
		createHorizontalCorridor(grid, startX, endX, endY)
	}
}

// Функция для создания горизонтального коридора
func createHorizontalCorridor(grid [][]models.Cell, startX, endX, y int) {
	// Определяем начало и конец коридора
	start := startX
	end := endX

	if startX > endX {
		start = endX
		end = startX
	}

	// Создаем горизонтальный коридор
	for x := start; x <= end; x++ {
		// Проверка границ сетки
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			grid[x][y].Filler = models.FillerMap["0"]
		}
	}
}

// Функция для создания вертикального коридора
func createVerticalCorridor(grid [][]models.Cell, startY, endY, x int) {
	// Определяем начало и конец коридора
	start := startY
	end := endY

	if startY > endY {
		start = endY
		end = startY
	}

	// Создаем вертикальный коридор
	for y := start; y <= end; y++ {
		// Проверка границ сетки
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			grid[x][y].Filler = models.FillerMap["0"]
		}
	}
}
