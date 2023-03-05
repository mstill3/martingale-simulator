package main

import "math/rand"

func getRandomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
