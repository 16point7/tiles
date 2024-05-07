package tiles

import "math/rand"

type game struct {
	board [][]uint
	score uint
}

func (g *game) MoveDown() (gameOver bool) {
	score, moved := moveDown(g.board)
	g.score += score
	if moved {
		newTile(g.board)
		gameOver = isTerminal(g.board)
	}
	return
}

func moveDown(board [][]uint) (score uint, moved bool) {
	for i := 0; i < length; i++ {
		lastCapturedJ := -1
		for j := length - 1; j > -1; j-- {
			nextJ := nextDown(board, j, i, lastCapturedJ)
			if nextJ == j {
				continue
			}
			if board[nextJ][i] == board[j][i] {
				lastCapturedJ = nextJ
				score += board[nextJ][i] * 2
				board[nextJ][i], board[j][i] = board[nextJ][i]*2, 0
			} else {
				board[nextJ][i], board[j][i] = board[j][i], 0
			}
			moved = true
		}
	}
	return
}

func nextDown(board [][]uint, j, i, lastCapturedJ int) int {
	nextJ := j + 1
	for nextJ < length && (board[nextJ][i] == 0 || j != lastCapturedJ && board[nextJ][i] == board[j][i]) {
		nextJ++
	}
	return nextJ - 1
}

func newTile(board [][]uint) {
	for {
		r := rand.Intn(area)
		j := r / length
		i := r % length

		if board[j][i] != 0 {
			continue
		}

		if rand.Float64() < 0.1 {
			board[j][i] = 4
		} else {
			board[j][i] = 2
		}

		break
	}
}

func isTerminal(board [][]uint) bool {
	for j := 1; j < length-1; j++ {
		for i := 1; i < length-1; i++ {
			if val := board[j][i]; val == board[j+1][i] || val == board[j-1][i] || val == board[j][i+1] || val == board[j][i-1] {
				return false
			}
		}
	}
	return true
}

func NewGame() *game {
	board := newBoard()
	initialize(board)
	return &game{
		board: board,
		score: 0,
	}
}

const (
	length = 4
	area   = length * length
)

func newBoard() [][]uint {
	temp := make([]uint, area)
	board := make([][]uint, length)
	for j := range board {
		board[j], temp = temp[:length], temp[length:]
	}
	return board
}

const numInitVals = 2

func initialize(board [][]uint) {
	for i := 0; i < numInitVals; i++ {
		newTile(board)
	}
}
