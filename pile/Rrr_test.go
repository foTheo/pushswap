package pile

import (
	"reflect"
	"testing"
)

func TestRrr(t *testing.T) {
	tests := []struct {
		name  string
		pile  *Pile
		wantA []int
		wantB []int
	}{
		{
			name:  "rotation inverse de a et b",
			pile:  &Pile{a: []int{1, 2, 3}, b: []int{4, 5, 6}},
			wantA: []int{3, 1, 2},
			wantB: []int{6, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rrr(tt.pile)
			if !reflect.DeepEqual(tt.pile.a, tt.wantA) {
				t.Errorf("Rrr() a = %v, want %v", tt.pile.a, tt.wantA)
			}
			if !reflect.DeepEqual(tt.pile.b, tt.wantB) {
				t.Errorf("Rrr() b = %v, want %v", tt.pile.b, tt.wantB)
			}
		})
	}
}
