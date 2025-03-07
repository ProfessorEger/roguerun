package dungeon_reader

import (
	"fmt"
	"os"
	"roguerun/models"
	"strings"
)

func ReadDungeonFromFile(filename string) [][][]models.Cell {
	var dungeon [][][]models.Cell

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return dungeon
	}
	dungeon = decodeDungeon(divideIntoWords(divideIntoLines(divideIntoParagraphs(string(data)))))

	return dungeon
}
func decodeDungeon(strDungeon [][][]string) [][][]models.Cell {
	var dungeon [][][]models.Cell = make([][][]models.Cell, len(strDungeon))
	for i := 0; i < len(dungeon); i++ {
		dungeon[i] = make([][]models.Cell, len(strDungeon[i]))
		for j := 0; j < len(dungeon[i]); j++ {
			dungeon[i][j] = make([]models.Cell, len(strDungeon[i][j]))
			for k := 0; k < len(dungeon[i][j]); k++ {
				dungeon[i][j][k] = translateStringToCell(strDungeon[i][j][k], [2]int{j, k})
			}
		}
	}
	return dungeon
}

func translateStringToCell(str string, coordinates [2]int) (cell models.Cell) {
	cell.Coordinates = coordinates
	cell.Filler = models.FillerMap[string(str[0])]
	cell.Creature = models.CreatureMap[string(str[1])]
	cell.Object = models.ObjectMap[string(str[2])]
	return
}

func divideIntoParagraphs(str string) []string {
	paragraphs := strings.Split(str, "\n\n")

	// Убираем возможные пустые строки по краям
	for i, paragraph := range paragraphs {
		paragraphs[i] = strings.TrimSpace(paragraph)
	}

	return paragraphs
}

func divideIntoLines(paragraphs []string) [][]string {
	var lines [][]string
	for i := 0; i < len(paragraphs); i++ {
		line := strings.Split(paragraphs[i], "\n")
		lines = append(lines, line)
	}
	return lines
}

func divideIntoWords(lines [][]string) [][][]string {
	var words [][][]string
	for i := 0; i < len(lines); i++ {
		var lineWords [][]string // Временный срез для хранения слов текущей строки
		for j := 0; j < len(lines[i]); j++ {
			// Разделяем строку на слова и добавляем в временный срез
			lineWords = append(lineWords, strings.Fields(lines[i][j]))
		}
		// Добавляем слова текущего абзаца в общий срез
		words = append(words, lineWords)
	}
	return words
}
