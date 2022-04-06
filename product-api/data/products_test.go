package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{Name: "Tst", Price: 1, SKU: "abcd-ef-a12"}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
