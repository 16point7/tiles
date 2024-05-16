package tiles

import "math/rand"

type game struct {
	Board [][]uint
	Score uint
}

func (g *game) MoveDown() (gameOver bool) {
	score, moved := moveDown(g.Board)
	g.Score += score
	if moved {
		newTile(g.Board)
		gameOver = isTerminal(g.Board)
	}
	return
}

func moveDown(board [][]uint) (score uint, moved bool) {
	for i := 0; i < Side; i++ {
		lastCapturedJ := -1
		for j := Side - 1; j > -1; j-- {
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
	if board[j][i] == 0 {
		return j
	}
	nextJ := j + 1
	for nextJ < Side && (board[nextJ][i] == 0 || j != lastCapturedJ && board[nextJ][i] == board[j][i]) {
		nextJ++
	}
	return nextJ - 1
}

func (g *game) MoveUp() (gameOver bool) {
	score, moved := moveUp(g.Board)
	g.Score += score
	if moved {
		newTile(g.Board)
		gameOver = isTerminal(g.Board)
	}
	return
}

func moveUp(board [][]uint) (score uint, moved bool) {
	for i := 0; i < Side; i++ {
		lastCapturedJ := -1
		for j := 0; j < Side; j++ {
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
	if board[j][i] == 0 {
		return j
	}
	nextJ := j - 1
	for nextJ > -1 && (board[nextJ][i] == 0 || j != lastCapturedJ && board[nextJ][i] == board[j][i]) {
		nextJ--
	}
	return nextJ + 1
}

func (g *game) MoveLeft() (gameOver bool) {
	score, moved := moveLeft(g.Board)
	g.Score += score
	if moved {
		newTile(g.Board)
		gameOver = isTerminal(g.Board)
	}
	return
}

func moveLeft(board [][]uint) (score uint, moved bool) {
	for j := 0; j < Side; j++ {
		lastCapturedJ := -1
		for i := 0; i < Side; i++ {
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
	if board[j][i] == 0 {
		return i
	}
	nextI := i - 1
	for nextI > -1 && (board[j][nextI] == 0 || j != lastCapturedJ && board[j][nextI] == board[j][i]) {
		nextI--
	}
	return nextI + 1
}

func (g *game) MoveRight() (gameOver bool) {
	score, moved := moveRight(g.Board)
	g.Score += score
	if moved {
		newTile(g.Board)
		gameOver = isTerminal(g.Board)
	}
	return
}

func moveRight(board [][]uint) (score uint, moved bool) {
	for j := 0; j < Side; j++ {
		lastCapturedJ := -1
		for i := Side - 1; i > -1; i-- {
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
	if board[j][i] == 0 {
		return i
	}
	nextI := i + 1
	for nextI < Side && (board[j][nextI] == 0 || j != lastCapturedJ && board[j][nextI] == board[j][i]) {
		nextI++
	}
	return nextI - 1
}

func newTile(board [][]uint) {
	vacancies := [Side * Side]struct {
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
	for j := 0; j < Side; j++ {
		for i := 0; i < Side-1; i++ {
			if val := board[j][i]; val == 0 || val == board[j][i+1] {
				return false
			}
		}
		if board[j][Side-1] == 0 {
			return false
		}
	}

	for i := 0; i < Side; i++ {
		for j := 0; j < Side-1; j++ {
			if val := board[j][i]; val == 0 || val == board[j+1][i] {
				return false
			}
		}
		if board[Side-1][i] == 0 {
			return false
		}
	}

	return true
}

func NewGame() *game {
	board := newBoard()
	initialize(board)
	return &game{
		Board: board,
		Score: 0,
	}
}

// Size of game board is Side X Side cells
const Side = 4

func newBoard() [][]uint {
	temp := make([]uint, Side*Side)
	board := make([][]uint, Side)
	for j := range board {
		board[j], temp = temp[:Side], temp[Side:]
	}
	return board
}

const numInitVals = 2

func initialize(board [][]uint) {
	for i := 0; i < numInitVals; i++ {
		newTile(board)
	}
}
