package kata

import "fmt"

type SnakesLadders struct {
	moveTo        map[int]int
	playerOne     int
	playerTwo     int
	complete      bool
	turnPlayerOne bool
}

func (sl *SnakesLadders) NewGame() {
	sl.moveTo = map[int]int{
		2:  38,
		7:  14,
		8:  31,
		15: 26,
		21: 42,
		28: 84,
		36: 44,
		51: 67,
		71: 91,
		78: 98,
		87: 94,
		16: 6,
		46: 25,
		49: 11,
		62: 19,
		64: 60,
		74: 53,
		89: 68,
		92: 88,
		95: 75,
		99: 80,
	}
	sl.playerOne = 0
	sl.playerTwo = 0
	sl.complete = false
	sl.turnPlayerOne = true
}

func (sl *SnakesLadders) Play(die1, die2 int) (r string) {
	if sl.complete {
		return "Game over!"
	}

	sum := die1 + die2

	if sl.turnPlayerOne {
		sl.playerOne += sum
		if sl.playerOne == 100 {
			sl.complete = true
			return "Player 1 Wins!"
		}
		if sl.playerOne > 100 {
			sl.playerOne = 200 - sl.playerOne
		}
		if v, ok := sl.moveTo[sl.playerOne]; ok {
			sl.playerOne = v
		}
		r = fmt.Sprintf("Player 1 is on square %d", sl.playerOne)
	} else {
		sl.playerTwo += sum
		if sl.playerTwo == 100 {
			sl.complete = true
			return "Player 2 Wins!"
		}
		if sl.playerTwo > 100 {
			sl.playerTwo = 200 - sl.playerTwo
		}
		if v, ok := sl.moveTo[sl.playerTwo]; ok {
			sl.playerTwo = v
		}
		r = fmt.Sprintf("Player 2 is on square %d", sl.playerTwo)
	}

	if die1 != die2 {
		sl.turnPlayerOne = !sl.turnPlayerOne
	}

	return r
}
