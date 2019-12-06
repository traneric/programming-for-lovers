package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Let's simulate an election!")
	// first, let's seed the PRNG
	rand.Seed(time.Now().UnixNano())

	// next, read in the electoral votes
	electoralVotes := ReadElectoralVotes("electoralVotes.txt")

	filename := "debates.txt"
	// read in polls
	polls := ReadPollingData(filename)

	numTrials := 100000
	marginOfError := 0.1

	probability1, probability2, probabilityTie := SimulateMultipleElections(polls, electoralVotes, numTrials, marginOfError)

	fmt.Println("Estimated probability of a candidate 1 win is", probability1)
	fmt.Println("Estimated probability of a candidate 2 win is", probability2)
	fmt.Println("Estimated probability of a tie is", probabilityTie)
}

//SimulateMultipleElections takes polling data as a map of state names to floats (for candidate 1), along with a 
//map of state names to electoral votes, a number of trials, and a margin of error in the polls.
//It returns three values: the estimated probabilities of candidate 1 winning, candidate 2 winning, and a tie.
func SimulateMultipleElections(polls map[string]float64, electoralVotes map[string]uint, numTrials int, marginOfError float64) (float64, float64, float64) {
	// keep track of number of simulated elections won by each candidate (and ties)
	winCount1 := 0
	winCount2 := 0
	tieCount := 0 // oh no!

	//simulate an election n times and update counts each time
	for i := 0; i < numTrials; i++ {
		//call SimulateOneElection as a subroutine
		votes1, votes2 := SimulateOneElection(polls, electoralVotes, marginOfError)
		// did candidate 1 or candidate 2 win?
		if votes1 > votes2 {
			winCount1++
		} else if votes1 < votes2 {
			winCount2++
		} else { //tie!
			tieCount++
		}
	}

	probability1 := float64(winCount1) / float64(numTrials)
	probability2 := float64(winCount2) / float64(numTrials)
	tieProbability := float64(tieCount) / float64(numTrials)
	return probability1, probability2, tieProbability
}

//SimulateOneElection takes a map of state names to polling percentages along with a map of state names to electoral college votes and a margin of error.
//It returns the number of EC votes for each of the two candidates in one simulated election.
func SimulateOneElection(polls map[string]float64, electoralVotes map[string]uint, marginOfError float64) (uint, uint) {
	var collegeVotes1 uint = 0
	var collegeVotes2 uint = 0

	// range over all the states, and simulate the election in each state.
	for state := range polls {
		poll := polls[state] // current polling value in the state for candidate 1
		numVotes := electoralVotes[state]
		// let's adjust polling value with some randomized "noise"
		adjustedPoll := AddNoise(poll, marginOfError)
		// now we check who won state based on the adjusted poll ...
		if adjustedPoll >= 0.5 {
			// candidate 1 wins! give them the EC votes for the state
			collegeVotes1 += numVotes
		} else {
			//candidate 2 wins!
			collegeVotes2 += numVotes
		}
	}

	return collegeVotes1, collegeVotes2
}

//AddNoise takes a polling value for candidate 1 and a margin of error. It returns an adjusted polling value for candidate 1 after adding random noise.
func AddNoise(poll float64, marginOfError float64) float64 {
	x := rand.NormFloat64() // random number from standard normal distribution
	x = x / 2               // x has ~95% chance of being between -1 and 1
	x = x * marginOfError   // x has 95% chance of being -marginOfError and marginOfError
	return x + poll
}
