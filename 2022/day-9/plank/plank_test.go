package plank

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		want    []Move
		wantErr bool
	}{
		{"single right", args{[]string{"R 1"}}, []Move{{R, 1}}, false},
		{"single left", args{[]string{"L 1"}}, []Move{{L, 1}}, false},
		{"single left", args{[]string{"Q 1"}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMove_applyToHMovement(t *testing.T) {
	type fields struct {
		dir   direction
		steps int
	}
	type args struct {
		f field
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		expectedH pos
	}{
		{"one step right", fields{R, 1}, args{f: InitialField()}, pos{1, 0}},
		{"5 step right", fields{R, 5}, args{f: InitialField()}, pos{5, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			move := Move{
				dir:   tt.fields.dir,
				steps: tt.fields.steps,
			}
			move.ApplyTo(tt.args.f)
			if !reflect.DeepEqual(tt.args.f.Pos[0], tt.expectedH) {
				t.Errorf("ApplyTo() = %v, want %v", tt.args.f.Pos[0], tt.expectedH)

			}
		})
	}
}

func TestMove_applyToTMovement(t *testing.T) {
	type fields struct {
		dir   direction
		steps int
	}
	type args struct {
		f field
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		expectedT      pos
		expectedVisits int
	}{
		{"one step right", fields{R, 1}, args{f: InitialField()}, pos{0, 0}, 1},
		{"5 step right", fields{R, 5}, args{f: InitialField()}, pos{4, 0}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			move := Move{
				dir:   tt.fields.dir,
				steps: tt.fields.steps,
			}
			move.ApplyTo(tt.args.f)
			if !reflect.DeepEqual(tt.args.f.Pos[len(tt.args.f.Pos)-1], tt.expectedT) {
				t.Errorf("ApplyTo() = %v, want %v", tt.args.f.Pos[len(tt.args.f.Pos)-1], tt.expectedT)
			}
			if count := tt.args.f.VisitedByTCount(); count != tt.expectedVisits {
				t.Errorf("VisitedByTCount() = %v, want %v", count, tt.expectedVisits)

			}
		})
	}
}

func (p pos) plus(o pos) pos {
	return pos{p.x + o.x, p.y + o.y}
}

func TestMove_applyTo10Knot(t *testing.T) {
	moves := []Move{{R, 5}, {U, 8}, {L, 8}}
	field := Initial10KnotField()
	applyMoves(moves, field)

	fmt.Println(field)
	paint(field, pos{-10, -10}, 30, 30)
}

func paint(field field, ll pos, width int, height int) {
	canvas := strings.Builder{}
	for i := height - 1; i >= 0; i-- {
		for j := 0; j < width; j++ {
			test := ll.plus(pos{j, i})
			index := -1
			for i, v := range field.Pos {
				if index == -1 && v == test {
					index = i
				}
			}
			switch index {
			case 0:
				canvas.WriteString("H")
			case -1:
				if test.x == 0 && test.y == 0 {
					canvas.WriteString("s")
				} else {
					canvas.WriteString(".")
				}
			default:
				canvas.WriteString(strconv.Itoa(index))
			}
		}
		canvas.WriteString("\n")
	}
	fmt.Print(canvas.String())
}

func applyMoves(m []Move, f field) {
	for _, v := range m {
		v.ApplyTo(f)
	}
}
