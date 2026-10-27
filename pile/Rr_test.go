package pile

import (
	"reflect"
	"testing"
)

func TestRr(t *testing.T) {
	tests := []struct {
		name  string
		pile  *Pile
		wantA []int
		wantB []int
	}{
		{
			name:  "rotation de a et b",
			pile:  &Pile{a: []int{1, 2, 3}, b: []int{4, 5, 6}},
			wantA: []int{2, 3, 1},
			wantB: []int{5, 6, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rr(tt.pile)
			if !reflect.DeepEqual(tt.pile.a, tt.wantA) {
				t.Errorf("Rr() a = %v, want %v", tt.pile.a, tt.wantA)
			}
			if !reflect.DeepEqual(tt.pile.b, tt.wantB) {
				t.Errorf("Rr() b = %v, want %v", tt.pile.b, tt.wantB)
			}
		})
	}
}
