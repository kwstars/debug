package baz

import (
	"testing"
)

func TestMyFunction(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test 1",
			arg:  "input",
			want: "Processed: input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyFunction(tt.arg); got != tt.want {
				t.Errorf("MyFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
