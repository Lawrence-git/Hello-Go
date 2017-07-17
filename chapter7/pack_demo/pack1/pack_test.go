package pack1

import "testing"

func TestReturnAnotherStr(t *testing.T) {
	if Pack1Int != 42 {
		t.Errorf(ReturnAnotherStr())
	}
}
