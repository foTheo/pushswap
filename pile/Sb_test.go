package pile

import (
	"reflect"
	"testing"
)

func TestSb(t *testing.T) {
	tests := []struct {
		name string
		pile *Pile
		want []int
	}{
		{
			name: "échange les deux premiers éléments de la pile b",
			pile: &Pile{a: []int{}, b: []int{2, 1, 3, 6, 8, 5}},
			want: []int{1, 2, 3, 6, 8, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sb(tt.pile)
			if !reflect.DeepEqual(tt.pile.b, tt.want) {
				t.Errorf("Sb() = %v, want %v", tt.pile.b, tt.want)
			}
		})
	}
}
