package maths

import "testing"

func TestDecimalDigits(t *testing.T) {
	f := 5.9874
	f2 := DecimalDigits(f,2)
	t.Log(f,f2)
}
