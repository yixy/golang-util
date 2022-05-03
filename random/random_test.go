package random

import "testing"

func TestRandomString(t *testing.T) {
	result := RandomString(32)
	if len(result) != 32 {
		t.Error(len(result))
	}
	t.Log(result)

}
