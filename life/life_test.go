package life

import "testing"

type args struct {
	x int
	y int
}

func TestField_NextVitality(t *testing.T) {
	const (
		x      = 1
		y      = 1
		height = 3
		width  = 3
	)

	tests := []struct {
		name  string
		cells [][]int
		want  int
	}{
		{
			name: "#1: Should return dead while less then two alive neighbors around the cell",
			cells: [][]int{
				{_alive, _dead, _dead},
				{_dead, _alive, _dead},
				{_dead, _dead, _dead},
			},
			want: _dead,
		},
		{
			name: "#2: Should return alive while two alive neighbors around the cell",
			cells: [][]int{
				{_alive, _dead, _dead},
				{_dead, _alive, _dead},
				{_dead, _dead, _alive},
			},
			want: _alive,
		},
		{
			name: "#3: Should return alive while tree alive neighbors around the cell",
			cells: [][]int{
				{_alive, _dead, _dead},
				{_dead, _alive, _dead},
				{_alive, _dead, _alive},
			},
			want: _alive,
		},
		{
			name: "#4: Should return dead while more then tree alive neighbors around the cell",
			cells: [][]int{
				{_alive, _dead, _alive},
				{_dead, _alive, _dead},
				{_alive, _dead, _alive},
			},
			want: _dead,
		},
		{
			name: "#5: Should return alive while tree alive neighbors around the cell when cell is dead",
			cells: [][]int{
				{_alive, _dead, _dead},
				{_dead, _dead, _dead},
				{_alive, _dead, _alive},
			},
			want: _alive,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				cells:  tt.cells,
				width:  width,
				height: height,
			}
			if got := f.NextVitality(x, y); got != tt.want {
				t.Errorf("NextVitality() = %v, want %v", got, tt.want)
			}
		})
	}
}
