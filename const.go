package main

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

const EMPTY_SYMBOL = "   "
