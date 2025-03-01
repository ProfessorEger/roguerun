package models

type Creature struct {
	IsCreature bool
	Symbol     string
}

type Object struct {
	IsObject bool
	Symbol   string
}

type Filler struct {
	Symbol string
	Empty  bool
}

type Cell struct {
	Coordinates [2]int
	Creature    Creature
	Object      Object
	Filler      Filler
	Symbol      string
}

var (
	MIN_LEAF_SIZE    int
	MIN_RECT_SIZE    int
	MAX_GRID_SIZE    int
	MIN_GRID_SIZE    int
	NUMBER_OF_FLOORS int
	EMPTY_SYMBOL     string
)

var FillerMap = map[string]Filler{
	"0": {Empty: true, Symbol: "   "},
	"1": {Empty: false, Symbol: " · "},
}

var CreatureMap = map[string]Creature{
	"0": {IsCreature: false, Symbol: "   "},
	"1": {IsCreature: true, Symbol: " M "},
	"2": {IsCreature: true, Symbol: " A "},
}

var ObjectMap = map[string]Object{
	"0": {IsObject: false, Symbol: "   "},
	"1": {IsObject: true, Symbol: " 1 "},
	"2": {IsObject: true, Symbol: " 2 "},
}
