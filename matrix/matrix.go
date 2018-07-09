package matrix

import (
	"strings"
	"strconv"
	"errors"
)

type Matrix struct {
	rowsCount, ColsCount int
	data                 [][]int
}

func New(str string) (*Matrix, error) {
	lines := strings.Split(str, "\n")
	m := Matrix{}
	for i, line := range lines {
		numbers := strings.Split(strings.TrimSpace(line), " ")
		if m.ColsCount == 0 {
			m.ColsCount = len(numbers)
		}
		if m.ColsCount != len(numbers) {
			return nil, errors.New("matrix must be rectangle")
		}
		m.data = append(m.data, make([]int, len(numbers)))
		for j, number := range numbers {
			x, err := strconv.Atoi(number)
			if err != nil {
				return nil, errors.New("not a number")
			}
			m.data[i][j] = x
		}
	}
	m.rowsCount = len(m.data)
	m.ColsCount = len(m.data[0])

	return &m, nil
}

func (m Matrix) Rows() [][]int {
	a := make([][]int, len(m.data))
	for rowNum := 0; rowNum < len(m.data); rowNum++ {
		a[rowNum] = make([]int, len(m.data[rowNum]))
		for colNum := 0; colNum < len(m.data[rowNum]); colNum++ {
			a[rowNum][colNum] = m.data[rowNum][colNum]
		}
	}
	return a
}

func (m Matrix) Cols() [][]int {
	a := make([][]int, len(m.data[0]))
	for rowNum := 0; rowNum < len(m.data[0]); rowNum++ {
		a[rowNum] = make([]int, len(m.data))
		for colNum := 0; colNum < len(m.data); colNum++ {
			a[rowNum][colNum] = m.data[colNum][rowNum]
		}
	}
	return a
}

func (m *Matrix) Set(rowNum, colNum, x int) bool {
	if (rowNum >= m.rowsCount) || (colNum >= m.ColsCount) || (rowNum < 0) || (colNum < 0) {
		return false
	}
	m.data[rowNum][colNum] = x
	return true
}
