package main

import (
	"fmt"
	"strconv"
)

type RouletteWheelTile struct {
	value string
	color string
}

func (tile RouletteWheelTile) String() string {
	return fmt.Sprintf("[%v %v]", tile.color, tile.value)
}

func (tile RouletteWheelTile) isEven() bool {
	num, err := strconv.Atoi(tile.value)
	if err != nil {
		fmt.Println("ERROR: Cannot convert tile value to integer")
		return false
	} else {
		return num%2 == 0
	}
}

func (tile RouletteWheelTile) isOdd() bool {
	num, err := strconv.Atoi(tile.value)
	if err != nil {
		fmt.Println("ERROR: Cannot convert tile value to integer")
		return false
	} else {
		return num%2 != 0
	}
}

func (tile RouletteWheelTile) isRed() bool {
	return tile.color == "red"
}

func (tile RouletteWheelTile) isBlack() bool {
	return tile.color == "black"
}

func (tile RouletteWheelTile) isGreen() bool {
	return tile.color == "green"
}
