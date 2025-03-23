package models

// TODO: could you write an Optional?
// type Optional[T any] struct {
// 	value T
// 	valid bool
// }

type Creature struct {
	IsCreature bool // TODO: Can creature not be a creature?
	Symbol     string
}

// TODO: if you don't want to use generics you can always just separate creature type
//       form type that can represent both abscence and presence of a creature.
//
// type Creature struct {
// 	Name string
// 	Age int
// }

// type OptionalCreature struct {
// 	Creature Creature
// 	IsThere bool
// }

// TODO: you could emulate an enum with something like this:

// type EntityKind int

// const (
// 	Creature EntityKind = iota
// 	Object
// 	Filler
// )

// You can then use this type to guarantee presence of only one of Creature, Filler, Object in your cell.

type Object struct {
	IsObject bool
	Symbol   string
}

type Filler struct { // TODO: if you use an enum filler could be included as two separate options, like: WALL, EMPTY.
	Symbol string // TODO: why would you duplicate in memory the same symbols many times over?
	Empty  bool
}

type Cell struct {
	Coordinates [2]int
	Creature    Creature
	Object      Object
	Filler      Filler
	Symbol      string // TODO: you shouldn't duplicate state, because that leads to state synchronization bugs
}

var ( // TODO: move to config reader? Maybe also separate config?
	MIN_LEAF_SIZE    int
	MIN_RECT_SIZE    int
	MAX_GRID_SIZE    int
	MIN_GRID_SIZE    int
	NUMBER_OF_FLOORS int
	EMPTY_SYMBOL     string
)

// TODO: why do you need global variables? You can make a struct and pass it everywhere where it is needed:
// type MapConfig struct {
// 	MinLeafSize    int
// 	MinRectSize    int
// 	MaxGridSize    int
// 	MinGridSize    int
// 	NumberOfFloors int
// 	EmptySymbol    string
// }

// TODO: maybe this can be a part of config?

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
