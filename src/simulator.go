package main

import (
	"fmt"
	"math"
)

type Simulator struct {
	totalAmountWillingToBet float64
	singleBetAmount         float64
	numRounds               int
	numGames                int
}

func (simulator Simulator) String() string {
	return fmt.Sprintf("Simulator [Total_to_Bet: %v, Single_Bet: %v, Rounds: %v, Games: %v]", simulator.totalAmountWillingToBet, simulator.singleBetAmount, simulator.numRounds, simulator.numGames)
}

func (simulator Simulator) Run() {

	// Output simulator variables to console
	fmt.Println("ROULETTE:")
	fmt.Println(fmt.Sprintf("\t- Number Games: %v", simulator.numGames))
	fmt.Println(fmt.Sprintf("\t- Number Rounds per Game: %v", simulator.numRounds))
	fmt.Println(fmt.Sprintf("\t- Total Amount Willing to Bet: %v", simulator.totalAmountWillingToBet))
	fmt.Println(fmt.Sprintf("\t- Single Bet Amount: %v", simulator.singleBetAmount))
	fmt.Println()

	// local variables
	rouletteWheel := getAmericanRouletteWheel()
	currentBetAmount := simulator.singleBetAmount
	amountChange := 0.0
	numGamesWon := 0
	totalAmountWon := 0.0
	numGamesEven := 0
	numGamesLost := 0
	totalAmountLost := 0.0

	// For each game...
	for game := 1; game <= simulator.numGames; game++ {
		fmt.Println(fmt.Sprintf("\tGame %v", game))

		// For each round...
		for round := 1; round <= simulator.numRounds; round++ {
			fmt.Println(fmt.Sprintf("\t\tRound %v", round))

			// if lost >= total amount willing to lose, stop playing
			if simulator.totalAmountWillingToBet+amountChange <= 0 {
				fmt.Println("\t\t\tHit Betting Cap, Unwilling to bet any more")
				break
			}

			// alternate between red and black color bet
			var betColor = ""
			if round%2 == 0 {
				betColor = "red"
			} else {
				betColor = "black"
			}
			fmt.Println(fmt.Sprintf("\t\t\t- Betting $%v on %s", currentBetAmount, betColor))

			// spin roulette wheel
			fmt.Println("\t\t\t- Spinning Wheel...")
			selectedTile := rouletteWheel.Spin()
			fmt.Println(fmt.Sprintf("\t\t\t- Wheel landed on: %v", selectedTile))

			// compare bet and actual tile
			if (betColor == "red" && selectedTile.isRed()) || (betColor == "black" && selectedTile.isBlack()) {
				// match
				fmt.Println("\t\t\t- WIN")
				amountChange += currentBetAmount
				fmt.Println(fmt.Sprintf("\t\t\t- Round Change: +$%v", currentBetAmount))
				currentBetAmount = simulator.singleBetAmount
			} else {
				// not a match
				fmt.Println("\t\t\t- LOSE")
				amountChange -= currentBetAmount
				fmt.Println(fmt.Sprintf("\t\t\t- Round Change: -$%v", math.Abs(currentBetAmount)))
				currentBetAmount *= 2
			}

			// output amount change in round
			if amountChange >= 0 {
				fmt.Println(fmt.Sprintf("\t\t\t- Total Change: +$%v", amountChange))
			} else {
				fmt.Println(fmt.Sprintf("\t\t\t- Total Change: -$%v", math.Abs(amountChange)))
			}
			fmt.Println(fmt.Sprintf("\t\t\t- Total: $%v", simulator.totalAmountWillingToBet+amountChange))
		}

		// output amount change in game
		if amountChange > 0 {
			fmt.Println(fmt.Sprintf("\t\tOutcome: Gained $%v", amountChange))
			numGamesWon += 1
			totalAmountWon += amountChange
		} else if amountChange < 0 {
			fmt.Println(fmt.Sprintf("\t\tOutcome: Lost $%v", math.Abs(amountChange)))
			numGamesLost += 1
			totalAmountLost += math.Abs(amountChange)
		} else {
			fmt.Println("\t\tOutcome: Broke even")
			numGamesEven += 1
		}

		// reset game variables
		currentBetAmount = simulator.singleBetAmount
		amountChange = 0.0
	}

	// Calculate statistics
	averageAmountWon := totalAmountWon / float64(simulator.numGames)
	averageAmountLost := totalAmountLost / float64(simulator.numGames)
	rateSuccess := float64(numGamesWon+numGamesEven) / float64(simulator.numGames) * 100

	//a := averageAmountWon / simulator.singleBetAmount / float64(simulator.numRounds)
	//b := averageAmountLost / simulator.singleBetAmount / float64(simulator.numRounds)
	d := averageAmountLost / averageAmountWon

	// Output overall game statistics to console
	fmt.Println("\nSTATISTICS:")
	fmt.Println(fmt.Sprintf("- Number Games Won: %v", numGamesWon))
	fmt.Println(fmt.Sprintf("- Average Amount Won: $%v", averageAmountWon))
	fmt.Println(fmt.Sprintf("- Number Games Break Even: %v", numGamesEven))
	fmt.Println(fmt.Sprintf("- Number Games Lost: %v", numGamesLost))
	fmt.Println(fmt.Sprintf("- Average Amount Lost: $%v", averageAmountLost))
	fmt.Println(fmt.Sprintf("- Rate of Success: %v%%", rateSuccess))
	//fmt.Println(fmt.Sprintf("- Win Amount to Bet Multiplier Rate: %v", a))
	//fmt.Println(fmt.Sprintf("- Lose Amount to Bet Multiplier Rate: %v", b))
	fmt.Println(fmt.Sprintf("- Risk Amount Rate: $1 dollar won - $%v dollar lost", d))
}
