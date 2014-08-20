package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ReadNextInt() (result int, err error) {
	var (
		parsedValue int64
		readValue   string
		readSize    int
	)

	readSize, err = fmt.Scan(&readValue)
	if err != nil {
		return
	} else if readSize < 1 {
		err = errors.New("Can't read input")
		return
	} else {
		parsedValue, err = strconv.ParseInt(readValue, 10, 64)
		if err != nil {
			return
		} else {
			result = int(parsedValue)
		}
	}

	return
}

func ReadNextLine() (result string, err error) {
	var (
		readSize int
	)

	readSize, err = fmt.Scan(&result)
	if readSize < 1 {
		err = errors.New("Can't read input")
	}

	return
}

type IGrid interface {
	Load() error
	Display()
	GetBoardSize() int
	GetBotPosition() (int, int)
	GetNextMove() string
	SetBoardSize(int)
	SetBotPosition(x, y int)
	FindPeach()
}

type Grid struct {
	BoardSize      int
	BotPositionX   int
	BotPositionY   int
	Grid           [][]string
	PeachPositionX int
	PeachPositionY int
}

func (g *Grid) Display() {
	fmt.Println("Grid:")
	fmt.Println("  Size:", g.BoardSize)
	fmt.Println("  Bot position:", g.BotPositionX, g.BotPositionY)
	fmt.Println("  Peach position:", g.PeachPositionX, g.PeachPositionY)
	fmt.Println("  Grid:", g.Grid)
}

func (g *Grid) FindPeach() {
	for i := 0; i < len(g.Grid); i++ {
		for j := 0; j < len(g.Grid[i]); j++ {
			if g.Grid[i][j] == "p" {
				g.PeachPositionX = i
				g.PeachPositionY = j
				return
			}
		}
	}
}

func (g *Grid) GetBoardSize() int {
	return g.BoardSize
}

func (g *Grid) GetBotPosition() (int, int) {
	return g.BotPositionX, g.BotPositionY
}

func (g *Grid) GetNextMove() string {
	var (
		xDiff, yDiff int
	)

	xDiff = int(math.Abs(float64(g.BotPositionX - g.PeachPositionX)))
	yDiff = int(math.Abs(float64(g.BotPositionY - g.PeachPositionY)))

	if xDiff > yDiff {
		// Move vertically
		if g.BotPositionX < g.PeachPositionX {
			return "DOWN"
		} else {
			return "UP"
		}
	} else {
		// Move horizontally
		if g.BotPositionY < g.PeachPositionY {
			return "RIGHT"
		} else {
			return "LEFT"
		}
	}
}

func (g *Grid) Load() (err error) {
	var (
		chars []string
		line  string
	)

	if g.BoardSize, err = ReadNextInt(); err != nil {
		return
	}

	if g.BotPositionX, err = ReadNextInt(); err != nil {
		return
	}

	if g.BotPositionY, err = ReadNextInt(); err != nil {
		return
	}

	g.Grid = make([][]string, g.BoardSize, g.BoardSize)

	for i := 0; i < g.BoardSize; i++ {
		if line, err = ReadNextLine(); err != nil {
			return
		}
		chars = strings.Split(line, "")
		g.Grid[i] = make([]string, g.BoardSize, g.BoardSize)
		for j := 0; j < g.BoardSize; j++ {
			g.Grid[i][j] = chars[j]
		}
	}

	return
}

func (g *Grid) SetBoardSize(newVal int) {
	g.BoardSize = newVal
}

func (g *Grid) SetBotPosition(newX, newY int) {
	g.BotPositionX = newX
	g.BotPositionY = newY
}

func main() {
	var (
		err      error
		grid     IGrid = new(Grid)
		nextMove string
	)

	if err = grid.Load(); err != nil {
		fmt.Println(err.Error())
	} else {
		grid.FindPeach()
		nextMove = grid.GetNextMove()
		//grid.Display()
		fmt.Println(nextMove)
	}
}
