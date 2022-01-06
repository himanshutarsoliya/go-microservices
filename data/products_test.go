package data

import "testing"

func TestCheckValidation(t *testing.T) {
	prod := Product{}
	err := prod.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
