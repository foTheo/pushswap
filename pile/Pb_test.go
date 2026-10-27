package pile

import (
	"reflect"
	"testing"
)

func TestPb(t *testing.T) {
	tests := []struct {
		name string
		pile *Pile
		want []int
	}{
		{
			name: "pousse a[0] sur b",
			pile: &Pile{a: []int{2, 1, 3}, b: []int{4, 7, 9}},
			want: []int{2, 4, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pb(tt.pile)
			if !reflect.DeepEqual(tt.pile.b, tt.want) {
				t.Errorf("Pb() = %v, want %v", tt.pile.b, tt.want)
			}
		})
	}
}
