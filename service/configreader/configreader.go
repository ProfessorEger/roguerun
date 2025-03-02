package configreader

import (
	"encoding/json"
	"os"
	"roguerun/models"
)

type Config struct {
	MIN_LEAF_SIZE,
	MIN_RECT_SIZE,
	MAX_GRID_SIZE,
	MIN_GRID_SIZE,
	NUMBER_OF_FLOORS int

	EMPTY_SYMBOL,
	WALL_SYMBOL string
}

func ApplyConfig(filename string) (err error) {
	var config *Config
	config, err = loadConstantsFromFile(filename)
	updateModels(config)

	return err
}

func loadConstantsFromFile(filename string) (config *Config, err error) {
	config = newConfig()
	file, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return
	}

	return
}

func updateModels(config *Config) {
	// TODO: global will go away?
	models.MIN_LEAF_SIZE = config.MIN_LEAF_SIZE
	models.MIN_RECT_SIZE = config.MIN_RECT_SIZE
	models.MAX_GRID_SIZE = config.MAX_GRID_SIZE
	models.MIN_GRID_SIZE = config.MIN_GRID_SIZE
	models.NUMBER_OF_FLOORS = config.NUMBER_OF_FLOORS
	models.EMPTY_SYMBOL = config.EMPTY_SYMBOL

	// Update FillerMap
	models.FillerMap["0"] = models.Filler{Empty: true, Symbol: config.EMPTY_SYMBOL}
	models.FillerMap["1"] = models.Filler{Empty: false, Symbol: config.WALL_SYMBOL}
}

func newConfig() *Config { // TODO: default config?
	return &Config{
		MIN_LEAF_SIZE:    4,
		MIN_RECT_SIZE:    3,
		MAX_GRID_SIZE:    29,
		MIN_GRID_SIZE:    12,
		NUMBER_OF_FLOORS: 3,

		EMPTY_SYMBOL: "   ",
		WALL_SYMBOL:  " · ",
	}
}
