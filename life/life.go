package life

import (
	"math/rand"
	"time"
)

const (
	_alive              = 1
	_dead               = 0
	_maxLivingNeighbors = 3
)

func GenerateField(width, height int) *Field {
	rand.Seed(time.Now().UnixNano())

	field := newField(width, height)
	for i := 0; i < (width * height / 2); i++ {
		field.setVitality(rand.Intn(width), rand.Intn(height), _alive)
	}

	return field
}

type Field struct {
	cells  [][]int
	width  int
	height int
}

func newField(width, height int) *Field {
	cells := make([][]int, height)
	for cols := range cells {
		cells[cols] = make([]int, width)
	}

	return &Field{cells: cells, width: width, height: height}
}

func (f *Field) setVitality(x, y int, vitality int) {
	x, y = f.getCellShift(x, y)

	f.cells[y][x] = vitality
}

func (f *Field) GetCells() [][]int {
	return f.cells
}

func (f *Field) getVitality(x, y int) int {
	x, y = f.getCellShift(x, y)

	return f.cells[y][x]
}

func (f *Field) getCellShift(x, y int) (int, int) {
	x += f.width
	x %= f.width
	y += f.height
	y %= f.height

	return x, y
}

func (f *Field) LivingNeighbors(x, y int) int {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if j != 0 || i != 0 && f.getVitality(x+i, y+j) > 0 {
				alive++
			}
		}
	}

	return alive
}

func (f *Field) NextVitality(x, y int) int {
	livingNeighbors := f.LivingNeighbors(x, y)
	isLiving := f.getVitality(x, y) > 0
	if livingNeighbors == _maxLivingNeighbors || livingNeighbors == _maxLivingNeighbors-1 && isLiving {
		return _alive
	}

	return _dead
}

func (f *Field) NextRound() *Field {
	field := newField(f.width, f.height)
	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			field.setVitality(x, y, f.NextVitality(x, y))
		}
	}

	return field
}
