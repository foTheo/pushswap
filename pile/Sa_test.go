package pile

import (
	"reflect"
	"testing"
)

func TestSa(t *testing.T) {
	tests := []struct {
		name string
		pile *Pile
		want []int
	}{
		{
			name: "échange les deux premiers éléments de la pile a",
			pile: &Pile{a: []int{2, 1, 3, 6, 8, 5}, b: []int{}},
			want: []int{1, 2, 3, 6, 8, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sa(tt.pile)
			if !reflect.DeepEqual(tt.pile.a, tt.want) {
				t.Errorf("Sa() = %v, want %v", tt.pile.a, tt.want)
			}
		})
	}
}
