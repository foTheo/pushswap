package pile

import (
	"reflect"
	"testing"
)

func TestPa(t *testing.T) {
	tests := []struct {
		name string
		pile *Pile
		want []int
	}{
		{
			name: "pousse b[0] sur a",
			pile: &Pile{a: []int{2, 1, 3}, b: []int{4, 7, 9}},
			want: []int{4, 2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pa(tt.pile)
			if !reflect.DeepEqual(tt.pile.a, tt.want) {
				t.Errorf("Pa() = %v, want %v", tt.pile.a, tt.want)
			}
		})
	}
}
