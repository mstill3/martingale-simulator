package main

import "fmt"

func main() {
	rouletteWheel := getAmericanRouletteWheel()
	selectedTile := rouletteWheel.Spin()
	fmt.Println(selectedTile)
}
