package design

import (
	"errors"
	"fmt"
)

type TicTacToe struct {
	board [][]string
	turn  string // 'X' or 'O'
	plays int

	coltotals   [][]int
	rowtotals   [][]int
	diag1totals []int
	diag2totals []int
}

const x = "X"
const o = "O"

func NewTicTacToe(size int) *TicTacToe {
	board := make([][]string, size)
	coltotals := make([][]int, size)
	rowtotals := make([][]int, size)
	diag1totals := make([]int, 2)
	diag2totals := make([]int, 2)
	for i := 0; i < size; i++ {
		board[i] = make([]string, size)
		coltotals[i] = make([]int, 2)
		rowtotals[i] = make([]int, 2)

	}

	return &TicTacToe{
		board: board,
		turn:  x,
		plays: 0,

		coltotals:   coltotals,
		rowtotals:   rowtotals,
		diag1totals: diag1totals,
		diag2totals: diag2totals,
	}
}

func (t *TicTacToe) Play(row int, col int) (bool, error) {

	if t.board[row][col] != "" {
		return false, errors.New("false move")
	}
	t.plays++
	t.board[row][col] = t.turn

	idx := 0
	if t.turn == o {
		idx = 1
	}

	t.coltotals[col][idx]++
	t.rowtotals[row][idx]++
	if row == col {
		t.diag1totals[idx]++
	}
	if row-col == len(t.board)-1 || col-row == len(t.board)-1 {
		t.diag2totals[idx]++
	}

	fmt.Printf("%v %v %v %v %v %v\n", t.board, t.turn, t.coltotals, t.rowtotals, t.diag1totals, t.diag2totals)

	if t.coltotals[col][idx] == len(t.board) ||
		t.rowtotals[row][idx] == len(t.board) ||
		t.diag1totals[idx] == len(t.board) ||
		t.diag2totals[idx] == len(t.board) {
		return true, nil
	}

	if t.turn == x {
		t.turn = o
	} else {
		t.turn = x
	}

	return false, nil
}

func (t *TicTacToe) isWin() bool {
	if t.plays < len(t.board)*len(t.board) { // off by 1?
		return false
	}
	//rowcou := 0
	//colcou := 0
	//
	//for i:=0; i< len(t.board); i++ {
	//
	//}
	return false
}

/*

	t := design.NewTicTacToe(2)
	res, err := t.Play(0, 0) //x
	fmt.Printf("result %v %v\n", res, err)
	res, err = t.Play(0,1) //0
	fmt.Printf("result %v %v\n", res, err)
	res, err = t.Play(1, 1) //x
	fmt.Printf("result %v %v\n", res, err)
	res, err = t.Play(0, 1) //o
	fmt.Printf("result %v %v\n", res, err)
	res, err = t.Play(1, 0) //o
	fmt.Printf("result %v %v\n", res, err)

*/
