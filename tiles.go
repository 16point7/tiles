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

func (g *game) MoveUp() (gameOver bool) {
	score, moved := moveUp(g.board)
	g.score += score
	if moved {
		newTile(g.board)
		gameOver = isTerminal(g.board)
	}
	return
}

func moveUp(board [][]uint) (score uint, moved bool) {
	for i := 0; i < length; i++ {
		lastCapturedJ := -1
		for j := 0; j < length; j++ {
			nextJ := nextUp(board, j, i, lastCapturedJ)
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

func nextUp(board [][]uint, j, i, lastCapturedJ int) int {
	nextJ := j - 1
	for nextJ > -1 && (board[nextJ][i] == 0 || j != lastCapturedJ && board[nextJ][i] == board[j][i]) {
		nextJ--
	}
	return nextJ + 1
}

func (g *game) MoveLeft() (gameOver bool) {
	score, moved := moveLeft(g.board)
	g.score += score
	if moved {
		newTile(g.board)
		gameOver = isTerminal(g.board)
	}
	return
}

func moveLeft(board [][]uint) (score uint, moved bool) {
	for j := 0; j < length; j++ {
		lastCapturedJ := -1
		for i := 0; i < length; i++ {
			nextI := nextLeft(board, j, i, lastCapturedJ)
			if nextI == i {
				continue
			}
			if board[j][nextI] == board[j][i] {
				lastCapturedJ = nextI
				score += board[j][nextI] * 2
				board[j][nextI], board[j][i] = board[j][nextI]*2, 0
			} else {
				board[j][nextI], board[j][i] = board[j][i], 0
			}
			moved = true
		}
	}
	return
}

func nextLeft(board [][]uint, j, i, lastCapturedJ int) int {
	nextI := i - 1
	for nextI > -1 && (board[j][nextI] == 0 || j != lastCapturedJ && board[j][nextI] == board[j][i]) {
		nextI--
	}
	return nextI + 1
}

func (g *game) MoveRight() (gameOver bool) {
	score, moved := moveRight(g.board)
	g.score += score
	if moved {
		newTile(g.board)
		gameOver = isTerminal(g.board)
	}
	return
}

func moveRight(board [][]uint) (score uint, moved bool) {
	for j := 0; j < length; j++ {
		lastCapturedJ := -1
		for i := length - 1; i > -1; i-- {
			nextI := nextRight(board, j, i, lastCapturedJ)
			if nextI == i {
				continue
			}
			if board[j][nextI] == board[j][i] {
				lastCapturedJ = nextI
				score += board[j][nextI] * 2
				board[j][nextI], board[j][i] = board[j][nextI]*2, 0
			} else {
				board[j][nextI], board[j][i] = board[j][i], 0
			}
			moved = true
		}
	}
	return
}

func nextRight(board [][]uint, j, i, lastCapturedJ int) int {
	nextI := i + 1
	for nextI < length && (board[j][nextI] == 0 || j != lastCapturedJ && board[j][nextI] == board[j][i]) {
		nextI++
	}
	return nextI - 1
}

func newTile(board [][]uint) {
	vacancies := [area]struct {
		j int
		i int
	}{}
	v := 0

	for j, row := range board {
		for i, val := range row {
			if val == 0 {
				vacancies[v].j = j
				vacancies[v].i = i
				v++
			}
		}
	}

	vacancy := vacancies[rand.Intn(v)]

	if rand.Float64() < 0.1 {
		board[vacancy.j][vacancy.i] = 4
	} else {
		board[vacancy.j][vacancy.i] = 2
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
