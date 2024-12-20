package main

import "testing"

//name, a, b, want, isErr

var tests = []struct {
	name string
	a float32
	b float32
	want float32
	isErr bool
}{
	{"valid", 10.0, 2.0, 5.0, false},
	{"invalide", 10.0, 0.0, 0.0, true},
}

func TestDivide(t *testing.T) {
	for _, tt := range tests {
		want, err := divide(tt.a, tt.b)
		if(err != nil) != tt.isErr {
			t.Errorf("%s: got error %v, want error %v", tt.name, err, tt.isErr)
		}
		if want != tt.want {
			t.Errorf("%s: got %v, want %v", tt.name, want, tt.want ) 
			
		}
	}
}