package kata

import "testing"

type testCase struct {
	die1, die2 int
	expected   string
}

var testCases = []testCase{
	{1, 1, "Player 1 is on square 38"},
	{1, 5, "Player 1 is on square 44"},
	{6, 2, "Player 2 is on square 31"},
	{1, 1, "Player 1 is on square 25"},
}

func TestGame(t *testing.T) {
	var game SnakesLadders

	game.NewGame()

	for _, test := range testCases {
		if output := game.Play(test.die1, test.die2); output != test.expected {
			t.Errorf(`Expected
				<string>: %s
			to equal
				<string>: %s`, output, test.expected)
		}
	}
}
