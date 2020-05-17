package jsq

import "testing"

func TestBuildJSONStreamQuerier(t *testing.T) {
	jsq := BuildJSONStreamQuerier("SELECT * FROM STREAM")
	if jsq == nil {
		t.Error("Test failed")
	}
}

