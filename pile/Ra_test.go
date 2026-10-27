package pile

import (
	"reflect"
	"testing"
)

func TestRa(t *testing.T) {
	tests := []struct {
		name string
		pile *Pile
		want []int
	}{
		{
			name: "rotation de a",
			pile: &Pile{a: []int{2, 1, 3}, b: []int{4, 7, 9}},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ra(tt.pile)
			if !reflect.DeepEqual(tt.pile.a, tt.want) {
				t.Errorf("Ra() = %v, want %v", tt.pile.a, tt.want)
			}
		})
	}
}
