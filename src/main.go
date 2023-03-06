package main

func main() {
	// increasing num rounds, increases success rate
	simulator := Simulator{
		totalAmountWillingToBet: 500.00,
		singleBetAmount:         5.0,
		numRounds:               20,
		numGames:                5_000,
	}
	simulator.Run()
}
