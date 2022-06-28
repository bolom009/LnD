package life

import (
	"math"
	"math/rand"
	"time"
)

const (
	_alive              = 1
	_dead               = 0
	_maxLivingNeighbors = 3
)

var randGlider = [][]int{
	{
		1, 0, 0,
		0, 1, 1,
		1, 1, 0,
	},
	{
		0, 1, 0,
		0, 0, 1,
		1, 1, 1,
	},
	{
		0, 0, 1,
		1, 0, 1,
		0, 1, 1,
	},
}

// GenerateField make new random field
func GenerateField(width, height int) *Field {
	rand.Seed(time.Now().UnixNano())

	var (
		glider = randGlider[rand.Intn(len(randGlider)-1)]
		field  = newField(width, height)
		x      = int(math.Round(float64(width) / 3))
		y      = x
		i      = 0
	)

	for _, st := range glider {
		field.setVitality(x+i, y, st)
		if i == 2 {
			i = 0
			y++
			continue
		}

		i++
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

// GetCells return values of field cells
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

// LivingNeighbors return alive neighbors
func (f *Field) LivingNeighbors(x, y int) int {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.getVitality(x+i, y+j) > 0 {
				alive++
			}
		}
	}

	return alive
}

// NextVitality return next status of cell (alive or dead)
func (f *Field) NextVitality(x, y int) int {
	livingNeighbors := f.LivingNeighbors(x, y)
	isLiving := f.getVitality(x, y) > 0
	//1. Any live cell with fewer than two live neighbors dies as if caused by underpopulation.
	//2. Any live cell with two or three live neighbors lives on to the next generation.
	//3. Any live cell with more than three live neighbors dies, as if by overcrowding.
	//4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
	if livingNeighbors == _maxLivingNeighbors || livingNeighbors == _maxLivingNeighbors-1 && isLiving {
		return _alive
	}

	return _dead
}

// NextRound run new round for field
func (f *Field) NextRound() *Field {
	field := newField(f.width, f.height)
	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			field.setVitality(x, y, f.NextVitality(x, y))
		}
	}

	return field
}
