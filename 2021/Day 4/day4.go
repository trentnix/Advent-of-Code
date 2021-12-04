package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const boardSize = 5

type BingoSquare struct {
	Called bool
	Value  string
}

type BingoBoard struct {
	Values [boardSize][boardSize]BingoSquare
	winner bool
}

func NewBingoBoard(lines []string) *BingoBoard {
	if len(lines) < boardSize || boardSize <= 0 {
		// insufficient input or invalid board size
		return nil
	}

	board := new(BingoBoard)
	board.winner = false

	for i := 0; i < boardSize; i++ {
		values := strings.Fields(lines[i])
		for j := 0; j < boardSize; j++ {
			board.Values[i][j].Value = values[j]
			board.Values[i][j].Called = false
		}
	}

	return board
}

func (b *BingoBoard) IsWinner() bool {
	if b.winner {
		return true
	}

	for i := 0; i < boardSize; i++ {
		rowWins, columnWins := true, true
		for j := 0; j < boardSize; j++ {
			// check rows
			if b.Values[i][j].Called != true {
				rowWins = false
			}

			// check columns
			if b.Values[j][i].Called != true {
				columnWins = false
			}
		}

		if rowWins || columnWins {
			b.winner = true
			return true
		}
	}

	return false
}

func (b *BingoBoard) CalculateScore(number string) int {
	unmarked := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			value, error := strconv.Atoi(b.Values[i][j].Value)
			if error != nil {
				log.Fatal(error.Error())
			}

			if b.Values[i][j].Called == false {
				unmarked += value
			}
		}
	}

	winningNumber, error := strconv.Atoi(number)
	if error != nil {
		log.Fatal(error.Error())
	}

	return unmarked * winningNumber
}

func (b *BingoBoard) SetCalled(value string) bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b.Values[i][j].Value == value {
				b.Values[i][j].Called = true
				return true
			}
		}
	}

	return false
}

func (b *BingoBoard) Clear() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			b.Values[i][j].Value = ""
			b.Values[i][j].Called = false
		}
	}
}

func (b *BingoBoard) Print(w io.Writer) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b.Values[i][j].Called {
				fmt.Fprintf(w, "%3s!", b.Values[i][j].Value)
			} else {
				fmt.Fprintf(w, "%3s ", b.Values[i][j].Value)
			}
		}

		fmt.Fprintf(w, "\n")
	}

}

func main() {
	var inputFile string
	inputFile = os.Args[1]

	fileContents, err := ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	numLines := len(fileContents)
	if numLines <= 2 {
		log.Fatal(fmt.Errorf("invalid input data"))
	}

	numBoards := int((numLines - 1) / (int(boardSize) + 1))
	if numBoards <= 0 {
		log.Fatal(fmt.Errorf("no board data"))
	}

	calledNumbers := strings.Split(fileContents[0], ",")
	if len(calledNumbers) <= 0 {
		log.Fatal(fmt.Errorf("no called numbers data"))
	}

	var boards []*BingoBoard

	// import board data
	for i := 0; i < numBoards; i++ {
		startIndex := 2 + (boardSize+1)*i // first two lines are the called list and an empty line - each board ends with an empty line
		boardInput := fileContents[startIndex : startIndex+boardSize]

		boards = append(boards, NewBingoBoard(boardInput))
	}

	var firstWinningBoard, lastWinningBoard *BingoBoard
	var firstWinningNumber, lastWinningNumber string

	for _, calledNumber := range calledNumbers {
		for _, board := range boards {
			if !board.IsWinner() {
				board.SetCalled(calledNumber)

				if board.IsWinner() {
					lastWinningBoard = board
					lastWinningNumber = calledNumber

					if firstWinningBoard == nil {
						firstWinningBoard = board
						firstWinningNumber = calledNumber
					}
				}
			}
		}
	}

	fmt.Printf("First winning board: \n")
	firstWinningBoard.Print(os.Stdout)
	fmt.Printf("Last Number: %s, SCORE: %d\n\n", firstWinningNumber, firstWinningBoard.CalculateScore(firstWinningNumber))

	fmt.Printf("Last winning board: \n")
	lastWinningBoard.Print(os.Stdout)
	fmt.Printf("Last Number: %s, SCORE: %d\n\n", lastWinningNumber, lastWinningBoard.CalculateScore(lastWinningNumber))
}
