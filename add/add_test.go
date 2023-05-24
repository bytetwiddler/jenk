package add

import "testing"

func TestAdd(t *testing.T) {
	res, err := Add(1, 1)
	if res != 2 || err != nil {
		t.Fatalf("Add(1,1) does not equal 2 - %v %v", res, err)
	}
}
