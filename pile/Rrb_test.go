package pile

import (
	"reflect"
	"testing"
)

func TestRrb(t *testing.T) {
	tests := []struct {
		name string
		pile *Pile
		want []int
	}{
		{
			name: "rotation inverse de b",
			pile: &Pile{a: []int{1, 2, 3}, b: []int{4, 5, 6}},
			want: []int{6, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rrb(tt.pile)
			if !reflect.DeepEqual(tt.pile.b, tt.want) {
				t.Errorf("Rrb() = %v, want %v", tt.pile.b, tt.want)
			}
		})
	}
}
