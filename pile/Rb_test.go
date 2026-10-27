package pile

import (
	"reflect"
	"testing"
)

func TestRb(t *testing.T) {
	tests := []struct {
		name string
		pile *Pile
		want []int
	}{
		{
			name: "rotation de b",
			pile: &Pile{a: []int{1, 2, 3}, b: []int{4, 5, 6}},
			want: []int{5, 6, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rb(tt.pile)
			if !reflect.DeepEqual(tt.pile.b, tt.want) {
				t.Errorf("Rb() = %v, want %v", tt.pile.b, tt.want)
			}
		})
	}
}
