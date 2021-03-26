package leetcode

import (
	"reflect"
	"testing"
)

func Test_IsValid(t *testing.T) {

	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "example1",
			arg:  "((",
			want: false,
		},
		{
			name: "example2",
			arg:  "()[]{}",
			want: true,
		},
		{
			name: "example3",
			arg:  "(]",
			want: false,
		},
		{
			name: "example4",
			arg:  "([)]",
			want: false,
		},
	}

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T){
	// 		 if got := (); !reflect.DeepEqual(got, tt.want){
	// 			t.Errorf()
	// 			}
	// 	})
	// }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsValid = %v want = %v\n", got, tt.want)
			}
		})
	}

}
