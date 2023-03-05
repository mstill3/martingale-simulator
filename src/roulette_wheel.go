package main

import (
	"fmt"
)

type RouletteWheel struct {
	region string
	tiles  []RouletteWheelTile
}

func (wheel RouletteWheel) String() string {
	return fmt.Sprintf("%v Roulette Wheel", wheel.region)
}

func (wheel RouletteWheel) Spin() RouletteWheelTile {
	numTiles := len(wheel.tiles)
	if numTiles == 0 {
		fmt.Println("ERROR: No tiles in wheel")
		return RouletteWheelTile{"nil", "green"}
	} else {
		return wheel.tiles[getRandomInt(0, numTiles-1)]
	}
}

func getAmericanRouletteWheel() *RouletteWheel {
	const RED = "red"
	const GREEN = "green"
	const BLACK = "black"
	// counter-clockwise
	americanRouletteWheel := RouletteWheel{
		region: "America",
		tiles: []RouletteWheelTile{
			{"0", GREEN},
			{"2", BLACK},
			{"14", RED},
			{"35", BLACK},
			{"23", RED},
			{"4", BLACK},
			{"16", RED},
			{"33", BLACK},
			{"21", RED},
			{"6", BLACK},
			{"18", RED},
			{"31", BLACK},
			{"19", RED},
			{"8", BLACK},
			{"12", RED},
			{"29", BLACK},
			{"25", RED},
			{"10", BLACK},
			{"27", RED},
			{"00", GREEN},
			{"1", RED},
			{"13", BLACK},
			{"36", RED},
			{"24", BLACK},
			{"3", RED},
			{"15", BLACK},
			{"34", RED},
			{"22", BLACK},
			{"5", RED},
			{"17", BLACK},
			{"32", RED},
			{"20", BLACK},
			{"7", RED},
			{"11", BLACK},
			{"30", RED},
			{"26", BLACK},
			{"9", RED},
			{"28", BLACK},
		}}
	return &americanRouletteWheel
}
