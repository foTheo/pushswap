package pile

import (
	"reflect"
	"testing"
)

func TestRra(t *testing.T) {
	tests := []struct {
		name string
		pile *Pile
		want []int
	}{
		{
			name: "rotation inverse de a",
			pile: &Pile{a: []int{2, 1, 3}, b: []int{4, 7, 9}},
			want: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rra(tt.pile)
			if !reflect.DeepEqual(tt.pile.a, tt.want) {
				t.Errorf("Rra() = %v, want %v", tt.pile.a, tt.want)
			}
		})
	}
}
