package main

func main() {
	simulator := Simulator{
		totalAmountWillingToBet: 500.00,
		singleBetAmount:         5.0,
		// increasing num rounds, increases success rate
		numRounds: 20,
		numGames:  5_000,
	}
	simulator.Run()
}
