package pile

import (
	"reflect"
	"testing"
)

func TestSs(t *testing.T) {
	tests := []struct {
		name  string
		pile  *Pile
		wantA []int
		wantB []int
	}{
		{
			name:  "échange les deux premiers éléments de la pile a et b",
			pile:  &Pile{a: []int{2, 1, 3, 6, 8, 5}, b: []int{4, 7, 9}},
			wantA: []int{1, 2, 3, 6, 8, 5},
			wantB: []int{7, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ss(tt.pile)
			if !reflect.DeepEqual(tt.pile.a, tt.wantA) {
				t.Errorf("Ss() = %v, want %v", tt.pile.a, tt.wantA)
			}
			if !reflect.DeepEqual(tt.pile.b, tt.wantB) {
				t.Errorf("Ss() = %v, want %v", tt.pile.b, tt.wantB)
			}
		})
	}
}
